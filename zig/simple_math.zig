const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [50]u8 = undefined;

    std.debug.print("What is the first number? ", .{});
    var input = try stdin.readUntilDelimiter(buffer[0..], '\n');
    const first_number = try std.fmt.parseInt(i32, input, 10);

    std.debug.print("What is the second number? ", .{});
    input = try stdin.readUntilDelimiter(buffer[0..], '\n');
    const second_number = try std.fmt.parseInt(i32, input, 10);

    std.debug.print("{d} + {d} = {d}\n", .{ first_number, second_number, first_number + second_number });
    std.debug.print("{d} - {d} = {d}\n", .{ first_number, second_number, first_number - second_number });
    std.debug.print("{d} * {d} = {d}\n", .{ first_number, second_number, first_number * second_number });
    std.debug.print("{d} / {d} = {d}\n", .{ first_number, second_number, if (second_number <= 0) 0 else @divTrunc(first_number, second_number) });
}
