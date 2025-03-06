const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var quote_buffer: [200]u8 = undefined;

    std.debug.print("What is the quote? ", .{});
    const quote = try stdin.readUntilDelimiter(quote_buffer[0..], '\n');

    var author_buffer: [30]u8 = undefined;
    std.debug.print("Who said it? ", .{});
    const author = try stdin.readUntilDelimiter(author_buffer[0..], '\n');

    std.debug.print("{s} says, \"{s}\"", .{ author, quote });
}
