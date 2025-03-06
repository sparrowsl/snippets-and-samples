const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var noun_buffer: [30]u8 = undefined;
    var verb_buffer: [30]u8 = undefined;
    var adjective_buffer: [30]u8 = undefined;
    var adverb_buffer: [30]u8 = undefined;

    std.debug.print("Enter a noun: ", .{});
    const noun = try stdin.readUntilDelimiter(noun_buffer[0..], '\n');

    std.debug.print("Enter a verb: ", .{});
    const verb = try stdin.readUntilDelimiter(verb_buffer[0..], '\n');

    std.debug.print("Enter a adjective: ", .{});
    const adjective = try stdin.readUntilDelimiter(adjective_buffer[0..], '\n');

    std.debug.print("Enter a adverb: ", .{});
    const adverb = try stdin.readUntilDelimiter(adverb_buffer[0..], '\n');

    std.debug.print("Do you {s} your {s} {s} {s}? That's hilarious!", .{ verb, adjective, noun, adverb });
}
