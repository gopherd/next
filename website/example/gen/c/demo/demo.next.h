/* Code generated by "next 0.0.4"; DO NOT EDIT. */

#ifndef DEMO_DEMO_H
#define DEMO_DEMO_H

#include <stdint.h>
#include <string.h>

#if !defined(__cplusplus) && !defined(__bool_true_false_are_defined)
#include <stdbool.h>
#endif

#if !defined(_Bool)
typedef unsigned char _Bool;
#endif

// Enums forward declarations
typedef enum DEMO_Color DEMO_Color;
typedef /* enum */ struct DEMO_MathConstants DEMO_MathConstants;
typedef /* enum */ struct DEMO_OperatingSystem DEMO_OperatingSystem;

// Structs forward declarations
typedef struct DEMO_User DEMO_User;
typedef struct DEMO_Uint64 DEMO_Uint64;
typedef struct DEMO_Uint128 DEMO_Uint128;
typedef struct DEMO_Contract DEMO_Contract;
typedef struct DEMO_LoginRequest DEMO_LoginRequest;
typedef struct DEMO_LoginResponse DEMO_LoginResponse;

#define DEMO_VERSION "1.0.0" /* String constant */
#define DEMO_MAX_RETRIES 3 /* Integer constant */
#define DEMO_TIMEOUT 3000.0F /* Float constant expression */

/**
 * Color represents different color options
 * Values: Red (1), Green (2), Blue (4), Yellow (8)
 */
typedef enum DEMO_Color {
	DEMO_Color_Red = 1,
	DEMO_Color_Green = 2,
	DEMO_Color_Blue = 4,
	DEMO_Color_Yellow = 8,
} DEMO_Color;

/**
 * MathConstants represents mathematical constants
 */
typedef struct DEMO_MathConstants {
	double value;
} DEMO_MathConstants;

#define DEMO_MathConstants_Pi ((DEMO_MathConstants){3.14159265358979323846})
#define DEMO_MathConstants_E ((DEMO_MathConstants){2.71828182845904523536})

/**
 * OperatingSystem represents different operating systems
 */
typedef struct DEMO_OperatingSystem {
	const char* value;
} DEMO_OperatingSystem;

#define DEMO_OperatingSystem_Windows ((DEMO_OperatingSystem){"windows"})
#define DEMO_OperatingSystem_Linux ((DEMO_OperatingSystem){"linux"})
#define DEMO_OperatingSystem_MacOS ((DEMO_OperatingSystem){"macos"})
#define DEMO_OperatingSystem_Android ((DEMO_OperatingSystem){"android"})
#define DEMO_OperatingSystem_IOS ((DEMO_OperatingSystem){"ios"})

/**
 * User represents a user in the system
 */
typedef struct DEMO_User {
	int64_t id;
	char* username;
	void* tags;
	void* scores;
	double coordinates[3];
	int32_t matrix[3][2];
	char* email;
	DEMO_Color favorite_color;
	/**
	 * @next(tokens) applies to the node name:
	 * - For snake_case: "last_login_ip"
	 * - For camelCase: "lastLoginIP"
	 * - For PascalCase: "LastLoginIP"
	 * - For kebab-case: "last-login-ip"
	 */
	char* last_login_ip;
	void* extra;
} DEMO_User;

/**
 * uint64 represents a 64-bit unsigned integer.
 * - In Go, it is aliased as uint64
 * - In C++, it is aliased as uint64_t
 * - In Java, it is aliased as long
 * - In Rust, it is aliased as u64
 * - In C#, it is aliased as ulong
 * - In Protobuf, it is represented as uint64
 * - In other languages, it is represented as a struct with low and high 32-bit integers.
 */
typedef struct DEMO_Uint64 {
	int32_t low;
	int32_t high;
} DEMO_Uint64;

/**
 * uint128 represents a 128-bit unsigned integer.
 * - In rust, it is aliased as u128
 * - In other languages, it is represented as a struct with low and high 64-bit integers.
 */
typedef struct DEMO_Uint128 {
	DEMO_Uint64 low;
	DEMO_Uint64 high;
} DEMO_Uint128;

/**
 * Contract represents a smart contract
 */
typedef struct DEMO_Contract {
	DEMO_Uint128 address;
	void* data;
} DEMO_Contract;

/**
 * LoginRequest represents a login request message (type 101)
 * @message annotation is a custom annotation that generates message types.
 */
typedef struct DEMO_LoginRequest {
	char* username;
	char* password;
	/**
	 * @next(optional) annotation is a builtin annotation that marks a field as optional.
	 */
	char* device;
	DEMO_OperatingSystem os;
	int64_t timestamp;
} DEMO_LoginRequest;

static inline int DEMO_LoginRequest_message_type() { return 101; }

/**
 * LoginResponse represents a login response message (type 102)
 */
typedef struct DEMO_LoginResponse {
	char* token;
	DEMO_User user;
} DEMO_LoginResponse;

static inline int DEMO_LoginResponse_message_type() { return 102; }

/**
 * Reader provides reading functionality
 */
typedef struct DEMO_Reader DEMO_Reader;

struct DEMO_Reader {
	void* context;
	int32_t (*read)(DEMO_Reader* self, uint8_t* buffer);
};

/**
 * @next(error) applies to the method:
 * - For Go: The method may return an error
 * - For C++/Java: The method throws an exception
 *
 * @next(mut) applies to the method:
 * - For C++: The method is non-const
 * - For other languages: This annotation may not have a direct effect
 *
 * @next(mut) applies to the parameter buffer:
 * - For C++: The parameter is non-const, allowing modification
 * - For other languages: This annotation may not have a direct effect,
 *   but indicates that the buffer content may be modified
 */
static inline int32_t DEMO_Reader_read(DEMO_Reader* self, uint8_t* buffer) {
	return self->read(self, buffer);
}

/**
 * HTTPClient provides HTTP request functionality
 */
typedef struct DEMO_HTTPClient DEMO_HTTPClient;

struct DEMO_HTTPClient {
	void* context;
	char* (*request)(DEMO_HTTPClient* self, char* url, char* method, char* body);
	char* (*request2)(DEMO_HTTPClient* self, char* url, char* method, char* body);
};

/**
 * Available for all languages
 */
static inline char* DEMO_HTTPClient_request(DEMO_HTTPClient* self, char* url, char* method, char* body) {
	return self->request(self, url, method, body);
}

static inline char* DEMO_HTTPClient_request2(DEMO_HTTPClient* self, char* url, char* method, char* body) {
	return self->request2(self, url, method, body);
}

#endif /* DEMO_DEMO_H */
