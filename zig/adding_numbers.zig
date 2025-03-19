const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [1024]u8 = undefined;

    std.debug.print("How many numbers to add? ", .{});
    var input = try stdin.readUntilDelimiter(buffer[0..], '\n');

    const numbers_limit: u32 = try std.fmt.parseInt(u32, input, 10) catch 0;
    var sum: i32 = 0;

    for (0..numbers_limit) |_| {
        std.debug.print("Enter a number: ", .{});
        input = try stdin.readUntilDelimiter(buffer[0..], '\n');

        const value = std.fmt.parseInt(i32, input, 10) catch 0;
        sum += value;
    }

    std.debug.print("The total is: {d}", .{sum});
}
