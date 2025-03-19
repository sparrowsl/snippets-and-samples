// FIX: a bug currently!!
const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [1024]u8 = undefined;

    std.debug.print("Enter two strings and I'll tell you if they are anagrams:\n", .{});
    std.debug.print("Enter the first string: ", .{});
    var input = try stdin.readUntilDelimiter(buffer[0..], '\n');
    const first_string = input;

    std.debug.print("Enter the second string: ", .{});
    input = try stdin.readUntilDelimiter(buffer[0..], '\n');
    const second_string = input;

    const anagram = is_anagram(first_string, second_string);
    std.debug.print("{any}\n", .{anagram});
    if (anagram) {
        std.debug.print("\"{s}\" and \"{s}\" are anagrams", .{ first_string, second_string });
    } else {
        std.debug.print("\"{s}\" and \"{s}\" are not anagrams", .{ first_string, second_string });
    }
}

fn is_anagram(first: []const u8, second: []const u8) bool {
    if (first.len != second.len) return false;

    for (first, 0..) |value, idx| {
        if (value == second[idx]) return true;
    }

    return false;
}
