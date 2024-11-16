# 单元测试

针对某个功能（函数）来验证功能的正确性

单元测试符合条件：编写一个单元测试的函数来进行目标的函数的测试（Sum）
+ 单元测试函数需要放到一个包里，独立一个测试包
  + 目标包（unittest.Sum) 
  + 单元测试包（unittest_test.TestSum)，目标测试函数的名称TestSum
+ 默认规则：单元测试包的名称：{targetpkg}_test 命名，同一个目录下，可以独立存在这个包的测试包
  + 测试包的文件名称: {filename}_test.go
  + 测试包的名称：{filename}_test
  + 函数签名：Test{funcname}(t Testing.T)
