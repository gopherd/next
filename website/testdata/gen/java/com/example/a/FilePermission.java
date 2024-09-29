// Code generated by "next 0.0.4"; DO NOT EDIT.

package com.example.a;

/**
 * Enum with complex iota usage
 */
public enum FilePermission {
    None((int) 0),
    Execute((int) 1),
    Write((int) 2),
    Read((int) 4),
    UserRead((int) 4),
    UserWrite((int) 32),
    UserExecute((int) 256),
    GroupRead((int) 2048),
    GroupWrite((int) 16384),
    GroupExecute((int) 131072),
    OthersRead((int) 1048576),
    OthersWrite((int) 8388608),
    OthersExecute((int) 67108864),
    /**
     * 4|32|256|2048|16384|131072|1048576|8388608|67108864
     * 4 + 32 + 256 + 2048 + 16384 + 131072 + 1048576 + 8388608 + 67108864
     */
    All((int) 76695844);

    private final int value;

    FilePermission(int value) {
        this.value = value;
    }

    public int getValue() {
        return value;
    }
}
