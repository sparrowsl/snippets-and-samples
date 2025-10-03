const std = @import("std");

pub fn main() !void {
    var buffer: [1024]u8 = undefined;
    var stdin = std.fs.File.stdin().reader(&buffer);
    const reader = &stdin.interface;

    const sum = count(reader);
    std.debug.print("sum: {d}\n", .{sum});
}

fn count(reader: *std.Io.Reader) i32 {
    const contents = reader.allocRemaining(std.heap.page_allocator, .unlimited) catch return 0;

    var it = std.mem.tokenizeAny(u8, contents, " \n\t\r");
    var wc: i32 = 0;

    while (it.next()) |_| {
        wc += 1;
    }

    return wc;
}

test "test count words" {
    var reader = std.Io.Reader.fixed("hello world\n");

    const s = count(&reader);
    std.debug.print("num: {d}\n", .{s});
    try std.testing.expectEqual(s, 2);
}
