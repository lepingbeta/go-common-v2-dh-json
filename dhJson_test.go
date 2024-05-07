/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 06:43:51
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 07:25:26
 * @FilePath     : /v2/go-common-v2-dh-json/dhJson_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhjson

import (
	"testing"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
)

func TestJsonEncodeIndent(t *testing.T) {
	// 测试用例
	tests := []struct {
		name     string      // 测试用例名称
		input    interface{} // 要编码的输入数据
		expected string      // 预期的JSON编码字符串
	}{
		{
			name:     "SimpleStruct",
			input:    struct{ Name string }{Name: "John"},
			expected: "{\n    \"Name\": \"John\"\n}",
		},
		{
			name:     "NestedStruct",
			input:    struct{ Details struct{ Name string } }{Details: struct{ Name string }{Name: "Jane"}},
			expected: "{\n    \"Details\": {\n        \"Name\": \"Jane\"\n    }\n}",
		},
		{
			name:     "EmptyStruct",
			input:    struct{}{},
			expected: "{}",
		},
		{
			name:     "StringInput",
			input:    "simple string",
			expected: "\"simple string\"",
		},
		{
			name:     "IntInput",
			input:    42,
			expected: "42",
		},
		// 可以添加更多测试用例以覆盖不同的输入类型和场景
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JsonEncodeIndent(tt.input)
			dhlog.Info(got)
			if got != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, got)
			}
			// 验证日志是否被调用，由于这是一个模拟函数，实际的验证可能需要使用模拟库如 testify 或 gomock
		})
	}
}
