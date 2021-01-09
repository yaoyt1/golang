## 从终端获取输入
1. fmt.Scanf(format string, a ...interface{}) 空格作为分隔符,遇到换行结束
2. fmt.Scan(a ...interface{})   空格和换行作为分隔符
3. fmt.Scanln(a ...interface{}) 空格作为分隔符， 遇到换行结束

## 从字符串中获取输入
1. fmt.Sscanf(source string ,format string, a ...interface{}) 空格作为分隔符
2. fmt.Sscan(source string ,a ...interface{})   空格和换行作为分隔符
3. fmt.Sscanln(source string ,a ...interface{}) 空格作为分隔符

## 从文件中获取输入
1. fmt.Fscanf(source string ,format string, a ...interface{}) 空格作为分隔符
2. fmt.Fscan(source string ,a ...interface{})   空格和换行作为分隔符
3. fmt.Fscanln(source string ,a ...interface{}) 空格作为分隔符

## 终端格式化输出
1. fmt.Printf(format string, a ...interface{}) 格式化输出并打印到终端
2. fmt.Println(a ...interface{})   把零个或者多个变量打印到终端,并换行
3. fmt.Print(a ...interface{})      把零个或者多个变量打印到终端

## 格式化返回一个字符串
1. fmt.Sprintf(format string, a ...interface{}) 格式化并返回字符串
2. fmt.Sprintln(a ...interface{})   把零个或者多个变量安装空格格式化换行,返回字符串
3. fmt.Sprint(a ...interface{})      把零个或者多个变量打印到终端

## 格式化输出到文件中
1. fmt.Fsprintf(format string, a ...interface{}) 格式化输出， 并写入文件
2. fmt.Fsprintln(a ...interface{})   把零个或者多个变量写入文件,并换行
3. fmt.Fsprint(a ...interface{})      把零个或者多个变量写入文件

