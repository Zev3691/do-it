package validata

// https://segmentfault.com/a/1190000040445612

// len：length 等于，长度相等
// max：小于等于
// min：大于等于
// eq：等于，字符串相等
// ne：不等于
// gt：大于
// gte：大于等于
// lt：小于
// lte：小于等于，例如lte=10；
// oneof：值中的一个，例如oneof=1 2

// 支持时间范围的比较lte
// 时间 RegTime time.Time `validate:"lte"` 小于等于当前时间

// 跨字段约束
// eqfield=ConfirmPassword
// eqcsfield=InnerStructField.Field

// 字符串规则
// contains=：包含参数子串
// containsany：包含参数中任意的 UNICODE 字符
// containsrune：包含参数表示的 rune 字符
// excludes：不包含参数子串
// excludesall：不包含参数中任意的 UNICODE 字符
// excludesrune：不包含参数表示的 rune 字符
// startswith：以参数子串为前缀
// endswith：以参数子串为后缀

// 使用unqiue来指定唯一性约束，对不同类型的处理如下：

// 对于数组和切片，unique约束没有重复的元素；
// 对于map，unique约束没有重复的值；
// 对于元素类型为结构体的切片，unique约束结构体对象的某个字段不重复，通过unqiue=name

// 特殊规则

// -：跳过该字段，不检验；
// |：使用多个约束，只需要满足其中一个，例如rgb|rgba；
// required：字段必须设置，不能为默认值；
// omitempty：如果字段未设置，则忽略它。
