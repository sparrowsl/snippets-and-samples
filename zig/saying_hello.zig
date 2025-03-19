const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [20]u8 = undefined;

    std.debug.print("What is your name? ", .{});
    const name = try stdin.readUntilDelimiter(&buffer, '\n');

    std.debug.print("Hello, {s}, nice to meet you!", .{name});
}
