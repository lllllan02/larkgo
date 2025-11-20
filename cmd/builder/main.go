package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		dir = flag.String("dir", "service", "要扫描的根目录（默认为 service）")
	)
	flag.Parse()

	if err := scanAndGenerate(*dir); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}

// scanAndGenerate 递归扫描目录，找到所有 types.go 文件并生成对应的 types_gen.go
func scanAndGenerate(rootDir string) error {
	var count int
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过 vendor 和隐藏目录
		if info.IsDir() {
			name := info.Name()
			if name == "vendor" || name == ".git" || name == "node_modules" || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			return nil
		}

		// 只处理 types.go 文件
		if filepath.Base(path) != "types.go" {
			return nil
		}

		dir := filepath.Dir(path)
		outputFile := filepath.Join(dir, "types_gen.go")

		fmt.Printf("🔍 发现 types.go: %s\n", path)

		if err := generateBuilder(path, outputFile); err != nil {
			fmt.Printf("   ⚠️  跳过（无需生成）: %v\n", err)
			return nil
		}

		count++
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("\n✅ 完成！共处理 %d 个文件\n", count)
	return nil
}

type generatorContext struct {
	methodsBuf     bytes.Buffer
	needStrconv    bool
	needFmt        bool
	needCore       bool
	hasContent     bool
	definedStructs map[string]bool
}

func generateBuilder(inputFile, outputFile string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, inputFile, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("解析文件失败: %w", err)
	}

	ctx := &generatorContext{
		definedStructs: make(map[string]bool),
	}

	// 1. 收集所有定义的结构体名称
	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if _, ok := typeSpec.Type.(*ast.StructType); ok {
				ctx.definedStructs[typeSpec.Name.Name] = true
			}
		}
		return true
	})

	// 2. 遍历所有类型定义并生成代码
	ast.Inspect(node, func(n ast.Node) bool {
		processStructDeclaration(n, ctx)
		return true
	})

	if !ctx.hasContent {
		return fmt.Errorf("没有找到需要生成 builder 方法的结构体")
	}

	return writeGeneratedFile(node.Name.Name, ctx, outputFile)
}

func processStructDeclaration(n ast.Node, ctx *generatorContext) {
	genDecl, ok := n.(*ast.GenDecl)
	if !ok {
		return
	}

	forceGen, skipGen := checkBuilderComments(genDecl)
	if skipGen {
		return
	}

	structType, structName := extractStructInfo(genDecl)
	if structType == nil {
		return
	}

	if !shouldGenerateStruct(structName, forceGen) {
		return
	}

	// 生成 New 函数
	generateNewFunction(ctx, structType, structName)

	// 生成 query 参数方法
	generateQueryMethods(ctx, structType, structName)

	// 生成 path 参数方法
	generatePathMethods(ctx, structType, structName)

	// 生成字段的 With 方法
	generateFieldMethods(ctx, structType, structName)
}

func checkBuilderComments(genDecl *ast.GenDecl) (forceGen, skipGen bool) {
	if genDecl.Doc == nil {
		return false, false
	}

	for _, comment := range genDecl.Doc.List {
		text := strings.TrimSpace(comment.Text)
		// 支持在标记后添加描述，如: //builder:gen 国际化群名称
		if strings.HasPrefix(text, "//builder:gen") {
			forceGen = true
		} else if strings.HasPrefix(text, "//builder:skip") {
			skipGen = true
		}
	}
	return
}

func extractStructInfo(genDecl *ast.GenDecl) (*ast.StructType, string) {
	if len(genDecl.Specs) == 0 {
		return nil, ""
	}

	typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, ""
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, ""
	}

	return structType, typeSpec.Name.Name
}

func shouldGenerateStruct(structName string, forceGen bool) bool {
	return forceGen || strings.HasSuffix(structName, "Req")
}

func generateNewFunction(ctx *generatorContext, structType *ast.StructType, structName string) {
	queryParamsField := findQueryParamsField(structType)
	pathParamsField := findPathParamsField(structType)
	hasQueryParams := queryParamsField != ""
	hasPathParams := pathParamsField != ""

	ctx.methodsBuf.WriteString(fmt.Sprintf("func New%s() *%s {\n", structName, structName))

	if hasQueryParams || hasPathParams {
		ctx.methodsBuf.WriteString(fmt.Sprintf("\treturn &%s{\n", structName))

		if hasQueryParams {
			queryParamsType := getQueryParamsType(structType)
			if strings.HasPrefix(queryParamsType, "core.") {
				ctx.needCore = true
			}
			ctx.methodsBuf.WriteString(fmt.Sprintf("\t\t%s: make(%s),\n", queryParamsField, queryParamsType))
		}

		if hasPathParams {
			pathParamsType := getPathParamsType(structType)
			if strings.HasPrefix(pathParamsType, "core.") {
				ctx.needCore = true
			}
			ctx.methodsBuf.WriteString(fmt.Sprintf("\t\t%s: make(%s),\n", pathParamsField, pathParamsType))
		}

		ctx.methodsBuf.WriteString("\t}\n")
	} else {
		ctx.methodsBuf.WriteString(fmt.Sprintf("\treturn &%s{}\n", structName))
	}

	ctx.methodsBuf.WriteString("}\n\n")
}

func generateQueryMethods(ctx *generatorContext, structType *ast.StructType, structName string) {
	queryParamsField := findQueryParamsField(structType)
	if queryParamsField == "" {
		return
	}

	var queryParams []queryParam
	for _, field := range structType.Fields.List {
		if isQueryParamsField(field) {
			queryParams = parseParamCommentsFromField(field)
			break
		}
	}

	for _, qp := range queryParams {
		ctx.hasContent = true

		methodName := toCamelCase(qp.key)
		paramName := toLowerCamelCase(qp.key)
		setterCode := generateParamSetterCode(ctx, queryParamsField, qp, paramName)

		ctx.methodsBuf.WriteString(fmt.Sprintf("func (req *%s) %s(%s %s) *%s {\n",
			structName, methodName, paramName, qp.typ, structName))
		ctx.methodsBuf.WriteString(setterCode)
		ctx.methodsBuf.WriteString("\treturn req\n")
		ctx.methodsBuf.WriteString("}\n\n")
	}
}

func generatePathMethods(ctx *generatorContext, structType *ast.StructType, structName string) {
	pathParamsField := findPathParamsField(structType)
	if pathParamsField == "" {
		return
	}

	var pathParams []queryParam
	for _, field := range structType.Fields.List {
		if isPathParamsField(field) {
			pathParams = parseParamCommentsFromField(field)
			break
		}
	}

	for _, pp := range pathParams {
		ctx.hasContent = true

		methodName := toCamelCase(pp.key)
		paramName := toLowerCamelCase(pp.key)
		setterCode := generateParamSetterCode(ctx, pathParamsField, pp, paramName)

		ctx.methodsBuf.WriteString(fmt.Sprintf("func (req *%s) %s(%s %s) *%s {\n",
			structName, methodName, paramName, pp.typ, structName))
		ctx.methodsBuf.WriteString(setterCode)
		ctx.methodsBuf.WriteString("\treturn req\n")
		ctx.methodsBuf.WriteString("}\n\n")
	}
}

// generateParamSetterCode 生成参数 setter 代码（通用于 query 和 path）
func generateParamSetterCode(ctx *generatorContext, paramsField string, param queryParam, paramName string) string {
	switch param.typ {
	case "string":
		return fmt.Sprintf("\treq.%s.Set(%q, %s)\n", paramsField, param.key, paramName)

	// 有符号整数类型
	case "int", "int8", "int16", "int32", "int64":
		ctx.needStrconv = true
		return fmt.Sprintf("\treq.%s.Set(%q, strconv.FormatInt(int64(%s), 10))\n", paramsField, param.key, paramName)

	// 无符号整数类型
	case "uint", "uint8", "uint16", "uint32", "uint64":
		ctx.needStrconv = true
		return fmt.Sprintf("\treq.%s.Set(%q, strconv.FormatUint(uint64(%s), 10))\n", paramsField, param.key, paramName)

	// 浮点数类型
	case "float32":
		ctx.needStrconv = true
		return fmt.Sprintf("\treq.%s.Set(%q, strconv.FormatFloat(float64(%s), 'f', -1, 32))\n", paramsField, param.key, paramName)
	case "float64":
		ctx.needStrconv = true
		return fmt.Sprintf("\treq.%s.Set(%q, strconv.FormatFloat(%s, 'f', -1, 64))\n", paramsField, param.key, paramName)

	// 布尔类型
	case "bool":
		ctx.needStrconv = true
		return fmt.Sprintf("\treq.%s.Set(%q, strconv.FormatBool(%s))\n", paramsField, param.key, paramName)

	// 其他类型使用 fmt.Sprintf
	default:
		ctx.needFmt = true
		return fmt.Sprintf("\treq.%s.Set(%q, fmt.Sprintf(\"%%v\", %s))\n", paramsField, param.key, paramName)
	}
}

func generateFieldMethods(ctx *generatorContext, structType *ast.StructType, structName string) {
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 || isQueryParamsField(field) || isPathParamsField(field) {
			continue
		}

		fieldName := field.Names[0].Name
		jsonTag := extractJsonTag(field)

		if jsonTag == "" || !ast.IsExported(fieldName) {
			continue
		}

		// 获取字段类型信息
		paramType, assignExpr := extractFieldTypeInfo(ctx, field, fieldName)
		if paramType == "" {
			continue
		}

		methodName := "With" + fieldName
		paramName := strings.ToLower(fieldName[:1]) + fieldName[1:]

		ctx.methodsBuf.WriteString(fmt.Sprintf("func (req *%s) %s(%s %s) *%s {\n",
			structName, methodName, paramName, paramType, structName))
		ctx.methodsBuf.WriteString(fmt.Sprintf("\t%s\n", assignExpr))
		ctx.methodsBuf.WriteString("\treturn req\n")
		ctx.methodsBuf.WriteString("}\n\n")

		ctx.hasContent = true
	}
}

func extractJsonTag(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}
	tagValue := strings.Trim(field.Tag.Value, "`")
	return parseTag(tagValue, "json")
}

// extractFieldTypeInfo 提取字段类型信息，返回 (参数类型, 赋值表达式)
func extractFieldTypeInfo(ctx *generatorContext, field *ast.Field, fieldName string) (paramType string, assignExpr string) {
	paramName := strings.ToLower(fieldName[:1]) + fieldName[1:]

	switch t := field.Type.(type) {
	case *ast.StarExpr:
		// 指针类型: *T
		typeName := extractTypeName(t.X)
		if typeName == "" {
			return "", ""
		}

		// 如果是指向结构体的指针，参数也使用指针类型，并直接赋值
		// 否则（如基本类型指针），参数使用值类型，赋值时取地址
		if ctx.definedStructs[typeName] {
			return "*" + typeName, fmt.Sprintf("req.%s = %s", fieldName, paramName)
		}
		return typeName, fmt.Sprintf("req.%s = &%s", fieldName, paramName)

	case *ast.ArrayType:
		// 切片/数组类型: []T -> 参数类型 ...T（可变参数），直接赋值
		typeName := extractTypeName(t.Elt)
		if typeName == "" {
			return "", ""
		}
		return "..." + typeName, fmt.Sprintf("req.%s = %s", fieldName, paramName)

	case *ast.MapType:
		// Map 类型: map[K]V -> 直接赋值
		keyType := extractTypeName(t.Key)
		valType := extractTypeName(t.Value)
		if keyType == "" || valType == "" {
			return "", ""
		}
		return fmt.Sprintf("map[%s]%s", keyType, valType), fmt.Sprintf("req.%s = %s", fieldName, paramName)

	case *ast.Ident:
		// 普通类型: T -> 直接赋值
		return t.Name, fmt.Sprintf("req.%s = %s", fieldName, paramName)

	case *ast.SelectorExpr:
		// 外部包类型: pkg.T -> 直接赋值
		typeName := extractTypeName(t)
		if typeName == "" {
			return "", ""
		}
		return typeName, fmt.Sprintf("req.%s = %s", fieldName, paramName)

	default:
		return "", ""
	}
}

// extractTypeName 提取类型名称
func extractTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return fmt.Sprintf("%s.%s", x.Name, t.Sel.Name)
		}
		return ""
	case *ast.StarExpr:
		// 嵌套指针 **T
		innerType := extractTypeName(t.X)
		if innerType == "" {
			return ""
		}
		return "*" + innerType
	case *ast.ArrayType:
		// 嵌套数组 [][]T
		elemType := extractTypeName(t.Elt)
		if elemType == "" {
			return ""
		}
		return "[]" + elemType
	default:
		return ""
	}
}

func writeGeneratedFile(packageName string, ctx *generatorContext, outputFile string) error {
	var buf bytes.Buffer
	buf.WriteString("// Code generated by cmd/builder. DO NOT EDIT.\n\n")
	buf.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	// 添加导入
	imports := buildImports(ctx)
	if imports != "" {
		buf.WriteString(imports)
	}

	// 添加生成的方法
	buf.Write(ctx.methodsBuf.Bytes())

	// 格式化代码
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("格式化失败，输出原始内容:")
		fmt.Println(buf.String())
		return fmt.Errorf("格式化代码失败: %w", err)
	}

	// 写入文件
	if err := os.WriteFile(outputFile, formatted, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	fmt.Printf("   ✅ 生成: %s\n", outputFile)
	return nil
}

func buildImports(ctx *generatorContext) string {
	if !ctx.needStrconv && !ctx.needFmt && !ctx.needCore {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString("import (\n")

	// 标准库
	hasStdLib := false
	if ctx.needFmt {
		buf.WriteString("\t\"fmt\"\n")
		hasStdLib = true
	}
	if ctx.needStrconv {
		buf.WriteString("\t\"strconv\"\n")
		hasStdLib = true
	}

	// 第三方库
	if ctx.needCore {
		if hasStdLib {
			buf.WriteString("\n")
		}
		buf.WriteString(fmt.Sprintf("\t\"%s/core\"\n", getModulePath()))
	}

	buf.WriteString(")\n\n")
	return buf.String()
}

// parseTag 解析 struct tag，返回指定 key 的值
func parseTag(tag, key string) string {
	// 简单的 tag 解析：查找 key:"value"
	keyPrefix := key + `:"`
	start := strings.Index(tag, keyPrefix)
	if start == -1 {
		return ""
	}
	start += len(keyPrefix)
	end := strings.Index(tag[start:], `"`)
	if end == -1 {
		return ""
	}
	value := tag[start : start+end]
	// 处理 json tag 的 omitempty 等后缀
	if comma := strings.Index(value, ","); comma != -1 {
		value = value[:comma]
	}
	return value
}

// findQueryParamsField 查找结构体中 QueryParams 类型的字段名
func findQueryParamsField(structType *ast.StructType) string {
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		// 检查字段类型是否是 QueryParams 或 core.QueryParams
		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Name == "QueryParams" {
				return field.Names[0].Name
			}
		case *ast.SelectorExpr:
			if t.Sel.Name == "QueryParams" {
				return field.Names[0].Name
			}
		}
	}
	return ""
}

// getQueryParamsType 获取 QueryParams 的完整类型名
func getQueryParamsType(structType *ast.StructType) string {
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		// 检查字段类型是否是 QueryParams 或 core.QueryParams
		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Name == "QueryParams" {
				return "QueryParams"
			}
		case *ast.SelectorExpr:
			if t.Sel.Name == "QueryParams" {
				if x, ok := t.X.(*ast.Ident); ok {
					return x.Name + ".QueryParams"
				}
			}
		}
	}
	return "QueryParams"
}

// findPathParamsField 查找结构体中 PathParams 类型的字段名
func findPathParamsField(structType *ast.StructType) string {
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		// 检查字段类型是否是 PathParams 或 core.PathParams
		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Name == "PathParams" {
				return field.Names[0].Name
			}
		case *ast.SelectorExpr:
			if t.Sel.Name == "PathParams" {
				return field.Names[0].Name
			}
		}
	}
	return ""
}

// getPathParamsType 获取 PathParams 的完整类型名
func getPathParamsType(structType *ast.StructType) string {
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		// 检查字段类型是否是 PathParams 或 core.PathParams
		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Name == "PathParams" {
				return "PathParams"
			}
		case *ast.SelectorExpr:
			if t.Sel.Name == "PathParams" {
				if x, ok := t.X.(*ast.Ident); ok {
					return x.Name + ".PathParams"
				}
			}
		}
	}
	return "PathParams"
}

// getModulePath 获取当前模块路径
func getModulePath() string {
	// 读取 go.mod 文件获取模块路径
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "github.com/lllllan02/larkgo" // 默认值
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimPrefix(line, "module ")
		}
	}
	return "github.com/lllllan02/larkgo" // 默认值
}

// queryParam 查询参数定义
type queryParam struct {
	key  string
	typ  string
	desc string
}

// parseParamCommentsFromField 从字段注释中解析参数（query/path）
func parseParamCommentsFromField(field *ast.Field) []queryParam {
	var params []queryParam

	// 检查字段的 Doc 注释
	if field.Doc != nil {
		for _, comment := range field.Doc.List {
			if param, ok := parseQueryComment(comment.Text); ok {
				params = append(params, param)
			}
		}
	}

	// 检查字段的 Comment 注释（行尾注释）
	if field.Comment != nil {
		for _, comment := range field.Comment.List {
			if param, ok := parseQueryComment(comment.Text); ok {
				params = append(params, param)
			}
		}
	}

	return params
}

// parseQueryComment 解析单个参数注释
// 格式：//@key(type):description
func parseQueryComment(text string) (queryParam, bool) {
	text = strings.TrimSpace(text)

	if !strings.HasPrefix(text, "//@") {
		return queryParam{}, false
	}

	text = strings.TrimPrefix(text, "//@")

	// 查找 ( 的位置
	openParen := strings.Index(text, "(")
	if openParen == -1 {
		return queryParam{}, false
	}

	key := text[:openParen]

	// 查找 ) 的位置
	closeParen := strings.Index(text[openParen:], ")")
	if closeParen == -1 {
		return queryParam{}, false
	}
	closeParen += openParen

	typ := text[openParen+1 : closeParen]

	// 查找 : 分隔描述
	var desc string
	if colonIdx := strings.Index(text[closeParen:], ":"); colonIdx != -1 {
		desc = strings.TrimSpace(text[closeParen+colonIdx+1:])
	}

	return queryParam{
		key:  key,
		typ:  typ,
		desc: desc,
	}, true
}

// isQueryParamsField 判断字段是否是 QueryParams 类型
func isQueryParamsField(field *ast.Field) bool {
	switch t := field.Type.(type) {
	case *ast.Ident:
		return t.Name == "QueryParams"
	case *ast.SelectorExpr:
		return t.Sel.Name == "QueryParams"
	}
	return false
}

// isPathParamsField 判断字段是否是 PathParams 类型
func isPathParamsField(field *ast.Field) bool {
	switch t := field.Type.(type) {
	case *ast.Ident:
		return t.Name == "PathParams"
	case *ast.SelectorExpr:
		return t.Sel.Name == "PathParams"
	}
	return false
}

// toCamelCase 转换为大驼峰命名 (user_id_type -> UserIdType)
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

// toLowerCamelCase 转换为小驼峰命名 (user_id_type -> userIdType)
func toLowerCamelCase(s string) string {
	camel := toCamelCase(s)
	if len(camel) > 0 {
		return strings.ToLower(camel[:1]) + camel[1:]
	}
	return camel
}
