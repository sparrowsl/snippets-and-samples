const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [255]u8 = undefined;

    std.debug.print("What is the quote? ", .{});
    const quote = try stdin.readUntilDelimiter(buffer[0..], '\n');

    std.debug.print("Who said it? ", .{});
    const author = try stdin.readUntilDelimiter(buffer[quote.len..], '\n');

    std.debug.print("{s} says, \"{s}\"", .{ author, quote });
}
