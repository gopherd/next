// Code generated by "next 0.0.4"; DO NOT EDIT.

#pragma once

#include <any>
#include <array>
#include <cstdint>
#include <map>
#include <string>
#include <unordered_map>
#include <vector>

namespace demo {
// Enums forward declarations
enum class Color;
/* enum */ class MathConstants;
/* enum */ class OperatingSystem;

// Classes forward declarations
class User;
class Uint128;
class Contract;
class LoginRequest;
class LoginResponse;

inline constexpr auto Version = "1.0.0"; // String constant
inline constexpr auto MaxRetries = 3; // Integer constant
inline constexpr auto Timeout = 3000.0F; // Float constant expression

// Color represents different color options
// Values: Red (1), Green (2), Blue (4), Yellow (8)
enum class Color : int8_t {
	Red = 1,
	Green = 2,
	Blue = 4,
	Yellow = 8,
};

/* enum */ class MathConstants {
private:
	double value;

public:
	static inline constexpr double Pi = 3.14159265358979323846;
	static inline constexpr double E = 2.71828182845904523536;

	MathConstants(const double& v) : value(v) {}

	bool operator==(const MathConstants& other) const {
		return value == other.value;
	}

	operator double() const {
		return value;
	}
};

/* enum */ class OperatingSystem {
private:
	std::string value;

public:
	static inline const std::string Windows = "windows";
	static inline const std::string Linux = "linux";
	static inline const std::string MacOS = "macos";
	static inline const std::string Android = "android";
	static inline const std::string IOS = "ios";

	OperatingSystem(const std::string& v) : value(v) {}

	bool operator==(const OperatingSystem& other) const {
		return value == other.value;
	}

	operator std::string() const {
		return value;
	}
};

// User represents a user in the system
class User {
public:
	int64_t id = {0};
	std::string username = {""};
	std::vector<std::string> tags;
	std::unordered_map<std::string, int> scores;
	std::array<double, 3> coordinates = {0.0};
	std::array<std::array<int, 2>, 3> matrix;
	std::string email = {""};
	Color favorite_color = {Color(0)};
	// @next(tokens) applies to the node name:
	// - For snake_case: "last_login_ip"
	// - For camelCase: "lastLoginIP"
	// - For PascalCase: "LastLoginIP"
	// - For kebab-case: "last-login-ip"
	std::string last_login_ip = {""};
	std::any extra;

public:
	User() = default;
	~User() = default;
};

// uint128 represents a 128-bit unsigned integer.
// - In rust, it is aliased as u128
// - In other languages, it is represented as a struct with low and high 64-bit integers.
class Uint128 {
public:
	uint64_t low;
	uint64_t high;

public:
	Uint128() = default;
	~Uint128() = default;
};

// Contract represents a smart contract
class Contract {
public:
	Uint128 address;
	std::any data;

public:
	Contract() = default;
	~Contract() = default;
};

// LoginRequest represents a login request message (type 101)
// @message annotation is a custom annotation that generates message types.
class LoginRequest {
public:
	std::string username = {""};
	std::string password = {""};
	// @next(optional) annotation is a builtin annotation that marks a field as optional.
	std::string device = {""};
	OperatingSystem os = {OperatingSystem("")};
	int64_t timestamp = {0};

	static inline int message_type() { return 101; }

public:
	LoginRequest() = default;
	~LoginRequest() = default;
};

// LoginResponse represents a login response message (type 102)
class LoginResponse {
public:
	std::string token = {""};
	User user;

	static inline int message_type() { return 102; }

public:
	LoginResponse() = default;
	~LoginResponse() = default;
};

// Reader provides reading functionality
class Reader {
public:
	virtual ~Reader() = default;
	// @next(error) applies to the method:
	// - For Go: The method may return an error
	// - For C++/Java: The method throws an exception
	//
	// @next(mut) applies to the method:
	// - For C++: The method is non-const
	// - For other languages: This annotation may not have a direct effect
	//
	// @next(mut) applies to the parameter buffer:
	// - For C++: The parameter is non-const, allowing modification
	// - For other languages: This annotation may not have a direct effect,
	//   but indicates that the buffer content may be modified
	virtual int read(std::vector<unsigned char>& buffer) = 0;
};

// HTTPClient provides HTTP request functionality
class HTTPClient {
public:
	virtual ~HTTPClient() = default;
	// Available for all languages
	virtual std::string request(const std::string& url, const std::string& method, const std::string& body) const = 0;
	virtual std::string request2(const std::string& url, const std::string& method, const std::string& body) const = 0;
	// Available for C++
	virtual std::string post(const std::string& url, const std::string& body) const = 0;
};
} // namespace demo