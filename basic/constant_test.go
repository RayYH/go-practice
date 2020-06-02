package basic

import (
    "fmt"
)

func ExampleShowPi() {
    fmt.Println("Pi:", Pi)
    // Output: Pi: 3.14159
}

func ExampleShowDays() {
    fmt.Println("Days:", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
    // Output: Days: 0 1 2 3 4 5 6
}

func ExampleShowPersonInfo() {
    fmt.Printf("%s's age is %d", name, age)
    // Output: Ray's age is 24
}

func ExampleShowIotaUsageOne() {
    fmt.Println("a b c:", a, b, c)
    // Output: a b c: 0 1 2
}

func ExampleShowIotaUsageTwo() {
    fmt.Println("d e f:", d, e, f)
    // Output: d e f: 0 1 2
}

func ExampleShowIotaUsageThree() {
    fmt.Println("g h i j k:", g, h, i, j, k)
    // Output: g h i j k: 0 1 string string 4
}

func ExampleShowIotaUsageFour() {
    fmt.Println("l m n o p:", l, m, n, o, p)
    // Output: l m n o p: 7 8 8 3 4
}

func ExampleShowIotaUsageFive() {
    fmt.Println("colors:", RED, ORANGE, YELLOW, INDIGO, VIOLET)
    // Output: colors: 0 1 2 5 6
}

func ExampleShowIotaUsageSix() {
    fmt.Println("size:", KB, MB, GB, TB, PB, EB, ZB, YB)
    // Output: size: 1024 1.048576e+06 1.073741824e+09 1.099511627776e+12 1.125899906842624e+15 1.152921504606847e+18 1.1805916207174113e+21 1.2089258196146292e+24
}

func ExampleDeclareConstInsideAFunc() {
    const Truth = true
    fmt.Println("Go?", Truth)
    // Output: Go? true
}
