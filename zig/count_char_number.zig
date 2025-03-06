const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [100]u8 = undefined;

    std.debug.print("What is the input string? ", .{});
    const string = try stdin.readUntilDelimiter(buffer[0..], '\n');

    if (string.len == 0) {
        std.debug.print("Please enter something..", .{});
    } else {
        std.debug.print("{s} has {d} characters.", .{ string, string.len });
    }
}
