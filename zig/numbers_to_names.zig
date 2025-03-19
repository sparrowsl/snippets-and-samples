const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [5]u8 = undefined;

    std.debug.print("Please enter the number of the month: ", .{});
    const input = try stdin.readUntilDelimiter(&buffer, '\n');

    const month_idx = try std.fmt.parseInt(i32, input, 10);

    const month_name = switch (month_idx) {
        1 => "January",
        2 => "February",
        3 => "March",
        4 => "April",
        5 => "May",
        6 => "June",
        7 => "July",
        8 => "August",
        9 => "September",
        10 => "October",
        11 => "November",
        12 => "December",
        else => "Invalid Month",
    };

    std.debug.print("The name of the month is {s}", .{month_name});
}
