// Code generated by "next v0.0.4"; DO NOT EDIT.

#pragma once

#include <any>
#include <array>
#include <cstdint>
#include <map>
#include <string>
#include <unordered_map>
#include <vector>

namespace demo::a {
// Enums forward declarations
enum class Color;
enum class FilePermission;
enum class Day;
enum class Month;
enum class IotatestEnum;

// Classes forward declarations
class Point2D;
class Point3D;
class Rectangle;
class ComplexStruct;
class User;
class UserProfile;
class LoginRequest;
class LoginResponse;
class GenericRequest;
class GenericResponse;

// XX constant
// XX value 2
inline constexpr auto XX = 1; // XX value
// Constants
inline constexpr auto ServerName = "Comprehensive Test Server";
inline constexpr auto Version = "1.0.0";
inline constexpr auto MaxConnections = 1000;
inline constexpr auto Pi = 3.14159265358979323846;
inline constexpr auto MaxInt64 = 9223372036854775807; // 2^63 - 1
inline constexpr auto MinInt64 = -9223372036854775808; // -2^63
// Constants with complex expressions
inline constexpr auto A = 1;
inline constexpr auto B = 3;
inline constexpr auto C = 9;
inline constexpr auto D = 7;
inline constexpr auto E = 28;
inline constexpr auto F = 1052;
inline constexpr auto G = 1052;
inline constexpr auto H = 5672;
inline constexpr auto I = 5673.618; // Approximation of golden ratio
inline constexpr auto J = 47.28015; // 120 is 5!
// Constants with function calls
inline constexpr auto StringLength = 13;
inline constexpr auto MinValue = 1;
inline constexpr auto MaxValue = 5673;
// Constants using built-in functions
inline constexpr auto IntFromBool = 1;
inline constexpr auto IntFromFloat = 3;
inline constexpr auto FloatFromInt = 42.0;
inline constexpr auto FloatFromBool = 0;
inline constexpr auto BoolFromInt = true;
inline constexpr auto BoolFromString = true;
inline constexpr auto FormattedString1 = "The answer is 42";
inline constexpr auto FormattedString2 = "Pi is approximately 3.14";
inline constexpr auto FormattedString3 = "Hello World\n";
// Constants for testing complex expressions and bitwise operations
inline constexpr auto Complex1 = 5673;
inline constexpr auto Complex2 = 78547;
inline constexpr auto Complex3 = 31;
inline constexpr auto Complex4 = 31;
inline constexpr auto Complex5 = 31;

// Enum with iota
enum class Color : int32_t {
    Red = 1,
    Green = 2,
    Blue = 4,
    Alpha = 8,
    Yellow = 3,
    Cyan = 6,
    Magenta = 5,
    White = 7,
};

// Enum with complex iota usage
enum class FilePermission : int32_t {
    None = 0,
    Execute = 1,
    Write = 2,
    Read = 4,
    UserRead = 4,
    UserWrite = 32,
    UserExecute = 256,
    GroupRead = 2048,
    GroupWrite = 16384,
    GroupExecute = 131072,
    OthersRead = 1048576,
    OthersWrite = 8388608,
    OthersExecute = 67108864,
    // 4|32|256|2048|16384|131072|1048576|8388608|67108864
    // 4 + 32 + 256 + 2048 + 16384 + 131072 + 1048576 + 8388608 + 67108864
    All = 76695844,
};

enum class Day : int32_t {
    Monday = 1,
    Tuesday = 2,
    Wednesday = 4,
    Thursday = 8,
    Friday = 16,
    Saturday = 32,
    Sunday = 64,
    Weekday = 31,
    Weekend = 96,
};

enum class Month : int32_t {
    January = 1,
    February = 2,
    March = 4,
    April = 8,
    May = 16,
    June = 32,
    July = 64,
    August = 128,
    September = 256,
    October = 512,
    November = 1024,
    December = 2048,
    Q1 = 7,
    Q2 = 56,
    Q3 = 448,
    Q4 = 3584,
};

// Test cases for iota
enum class IotatestEnum : int32_t {
    A = 0, // 0
    B = 1, // 1
    C = 0, // 0
    D = 2, // 2
    E = 0, // 0
    F = 1, // 1
    G = 0, // 0
};

// Struct types
class Point2D {
public:
    Point2D() = default;
    ~Point2D() = default;
    
    double x = {0};
    double y = {0};
};

class Point3D {
public:
    Point3D() = default;
    ~Point3D() = default;
    
    Point2D point;
    double z = {0};
};

class Rectangle {
public:
    Rectangle() = default;
    ~Rectangle() = default;
    
    Point2D topLeft;
    Point2D bottomRight;
};

// Struct with various field types
class ComplexStruct {
public:
    ComplexStruct() = default;
    ~ComplexStruct() = default;
    
    bool flag = {false};
    int8_t tinyInt = {0};
    int16_t smallInt = {0};
    int32_t mediumInt = {0};
    int64_t bigInt = {0LL};
    int defaultInt = {0};
    float singlePrecision = {0F};
    double doublePrecision = {0};
    std::string text = {""};
    unsigned char singleByte = {0};
    std::vector<unsigned char> byteArray;
    std::array<int,5> fixedArray = {0};
    std::vector<std::string> dynamicArray;
    std::vector<int> intArray;
    std::unordered_map<std::string,int> dictionary;
};

class User {
public:
    User() = default;
    ~User() = default;
    
    int64_t id = {0LL};
    std::string username = {""};
    std::string email = {""};
    Day preferredDay = {Day(0)};
    Month birthMonth = {Month(0)};
};

class UserProfile {
public:
    UserProfile() = default;
    ~UserProfile() = default;
    
    User user;
    std::string firstName = {""};
    std::string lastName = {""};
    int age = {0};
    std::vector<std::string> interests;
};

// message types
class LoginRequest {
public:
    LoginRequest() = default;
    ~LoginRequest() = default;
    
    std::string username = {""};
    std::string password = {""};
    std::string deviceId = {""};
    std::string twoFactorToken = {""};
};

class LoginResponse {
public:
    LoginResponse() = default;
    ~LoginResponse() = default;
    
    bool success = {false};
    std::string errorMessage = {""};
    std::string authenticationToken = {""};
    User user;
};

class GenericRequest {
public:
    GenericRequest() = default;
    ~GenericRequest() = default;
    
    std::string requestId = {""};
    int64_t timestamp = {0LL};
};

class GenericResponse {
public:
    GenericResponse() = default;
    ~GenericResponse() = default;
    
    std::string requestId = {""};
    int64_t timestamp = {0LL};
    bool success = {false};
    std::string errorCode = {""};
    std::string errorMessage = {""};
};

} // namespace demo::a