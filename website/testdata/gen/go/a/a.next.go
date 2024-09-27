// Code generated by "next v0.0.4"; DO NOT EDIT.

package a

import "strconv"

var _ = strconv.FormatInt
// XX constant
// XX value 2
const XX = 1 // XX value
// Constants
const ServerName = "Comprehensive Test Server"
const Version = "1.0.0"
const MaxConnections = 1000
const Pi = 3.14159265358979323846
const MaxInt64 = 9223372036854775807 // 2^63 - 1
const MinInt64 = -9223372036854775808 // -2^63
// Constants with complex expressions
const A = 1
const B = 3
const C = 9
const D = 7
const E = 28
const F = 1052
const G = 1052
const H = 5672
const I = 5673.618 // Approximation of golden ratio
const J = 47.28015 // 120 is 5!
// Constants with function calls
const StringLength = 13
const MinValue = 1
const MaxValue = 5673
// Constants using built-in functions
const IntFromBool = 1
const IntFromFloat = 3
const FloatFromInt = 42.0
const FloatFromBool = 0
const BoolFromInt = true
const BoolFromString = true
const FormattedString1 = "The answer is 42"
const FormattedString2 = "Pi is approximately 3.14"
const FormattedString3 = "Hello World\n"
// Constants for testing complex expressions and bitwise operations
const Complex1 = 5673
const Complex2 = 78547
const Complex3 = 31
const Complex4 = 31
const Complex5 = 31

// Enum with iota
type Color int32

const (
    ColorRed = 1
    ColorGreen = 2
    ColorBlue = 4
    ColorAlpha = 8
    ColorYellow = 3
    ColorCyan = 6
    ColorMagenta = 5
    ColorWhite = 7
)

func (x Color) String() string {
    switch x {
    case ColorRed:
        return "Red"
    case ColorGreen:
        return "Green"
    case ColorBlue:
        return "Blue"
    case ColorAlpha:
        return "Alpha"
    case ColorYellow:
        return "Yellow"
    case ColorCyan:
        return "Cyan"
    case ColorMagenta:
        return "Magenta"
    case ColorWhite:
        return "White"
    }
    return "Color(" + strconv.FormatInt(int64(x), 10) + ")"
}

// Enum with complex iota usage
type FilePermission int32

const (
    FilePermissionNone = 0
    FilePermissionExecute = 1
    FilePermissionWrite = 2
    FilePermissionRead = 4
    FilePermissionUserRead = 4
    FilePermissionUserWrite = 32
    FilePermissionUserExecute = 256
    FilePermissionGroupRead = 2048
    FilePermissionGroupWrite = 16384
    FilePermissionGroupExecute = 131072
    FilePermissionOthersRead = 1048576
    FilePermissionOthersWrite = 8388608
    FilePermissionOthersExecute = 67108864
    // 4|32|256|2048|16384|131072|1048576|8388608|67108864
    // 4 + 32 + 256 + 2048 + 16384 + 131072 + 1048576 + 8388608 + 67108864
    FilePermissionAll = 76695844
)

func (x FilePermission) String() string {
    switch x {
    case FilePermissionNone:
        return "None"
    case FilePermissionExecute:
        return "Execute"
    case FilePermissionWrite:
        return "Write"
    case FilePermissionRead:
        return "Read"
    case FilePermissionUserRead:
        return "UserRead"
    case FilePermissionUserWrite:
        return "UserWrite"
    case FilePermissionUserExecute:
        return "UserExecute"
    case FilePermissionGroupRead:
        return "GroupRead"
    case FilePermissionGroupWrite:
        return "GroupWrite"
    case FilePermissionGroupExecute:
        return "GroupExecute"
    case FilePermissionOthersRead:
        return "OthersRead"
    case FilePermissionOthersWrite:
        return "OthersWrite"
    case FilePermissionOthersExecute:
        return "OthersExecute"
    case FilePermissionAll:
        return "All"
    }
    return "FilePermission(" + strconv.FormatInt(int64(x), 10) + ")"
}

type Day int32

const (
    DayMonday = 1
    DayTuesday = 2
    DayWednesday = 4
    DayThursday = 8
    DayFriday = 16
    DaySaturday = 32
    DaySunday = 64
    DayWeekday = 31
    DayWeekend = 96
)

func (x Day) String() string {
    switch x {
    case DayMonday:
        return "Monday"
    case DayTuesday:
        return "Tuesday"
    case DayWednesday:
        return "Wednesday"
    case DayThursday:
        return "Thursday"
    case DayFriday:
        return "Friday"
    case DaySaturday:
        return "Saturday"
    case DaySunday:
        return "Sunday"
    case DayWeekday:
        return "Weekday"
    case DayWeekend:
        return "Weekend"
    }
    return "Day(" + strconv.FormatInt(int64(x), 10) + ")"
}

type Month int32

const (
    MonthJanuary = 1
    MonthFebruary = 2
    MonthMarch = 4
    MonthApril = 8
    MonthMay = 16
    MonthJune = 32
    MonthJuly = 64
    MonthAugust = 128
    MonthSeptember = 256
    MonthOctober = 512
    MonthNovember = 1024
    MonthDecember = 2048
    MonthQ1 = 7
    MonthQ2 = 56
    MonthQ3 = 448
    MonthQ4 = 3584
)

func (x Month) String() string {
    switch x {
    case MonthJanuary:
        return "January"
    case MonthFebruary:
        return "February"
    case MonthMarch:
        return "March"
    case MonthApril:
        return "April"
    case MonthMay:
        return "May"
    case MonthJune:
        return "June"
    case MonthJuly:
        return "July"
    case MonthAugust:
        return "August"
    case MonthSeptember:
        return "September"
    case MonthOctober:
        return "October"
    case MonthNovember:
        return "November"
    case MonthDecember:
        return "December"
    case MonthQ1:
        return "Q1"
    case MonthQ2:
        return "Q2"
    case MonthQ3:
        return "Q3"
    case MonthQ4:
        return "Q4"
    }
    return "Month(" + strconv.FormatInt(int64(x), 10) + ")"
}

// Test cases for iota
type IotatestEnum int32

const (
    IotatestEnumA = 0 // 0
    IotatestEnumB = 1 // 1
    IotatestEnumC = 0 // 0
    IotatestEnumD = 2 // 2
    IotatestEnumE = 0 // 0
    IotatestEnumF = 1 // 1
    IotatestEnumG = 0 // 0
)

func (x IotatestEnum) String() string {
    switch x {
    case IotatestEnumA:
        return "A"
    case IotatestEnumB:
        return "B"
    case IotatestEnumC:
        return "C"
    case IotatestEnumD:
        return "D"
    case IotatestEnumE:
        return "E"
    case IotatestEnumF:
        return "F"
    case IotatestEnumG:
        return "G"
    }
    return "IotatestEnum(" + strconv.FormatInt(int64(x), 10) + ")"
}

// Struct types
type Point2D struct {
    X float64
    Y float64
}

type Point3D struct {
    Point Point2D
    Z float64
}

type Rectangle struct {
    TopLeft Point2D
    BottomRight Point2D
}

// Struct with various field types
type ComplexStruct struct {
    Flag bool
    TinyInt int8
    SmallInt int16
    MediumInt int32
    BigInt int64
    DefaultInt int
    SinglePrecision float32
    DoublePrecision float64
    Text string
    SingleByte byte
    ByteArray []byte
    FixedArray [5]int
    DynamicArray []string
    IntArray []int
    Dictionary map[string]int
}

type User struct {
    Id int64
    Username string
    Email string
    PreferredDay Day
    BirthMonth Month
}

type UserProfile struct {
    User User
    FirstName string
    LastName string
    Age int
    Interests []string
}

// message types
type LoginRequest struct {
    Username string
    Password string
    DeviceId string
    TwoFactorToken string
}

type LoginResponse struct {
    Success bool
    ErrorMessage string
    AuthenticationToken string
    User User
}

type GenericRequest struct {
    RequestId string
    Timestamp int64
}

type GenericResponse struct {
    RequestId string
    Timestamp int64
    Success bool
    ErrorCode string
    ErrorMessage string
}
