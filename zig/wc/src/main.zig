const std = @import("std");

pub fn main() !void {
    var buffer: [1024]u8 = undefined;
    var stdin = std.fs.File.stdin().reader(&buffer);
    const reader = &stdin.interface;

    const sum = count(reader);
    std.debug.print("{d}", .{sum});
}

fn count(reader: *std.Io.Reader) i32 {
    const contents = reader.peekDelimiterExclusive('\n') catch return 0;

    var it = std.mem.tokenizeAny(u8, contents, " ");
    var wc: i32 = 0;

    while (it.next()) |_| {
        wc += 1;
    }

    return wc;
}

test "test count words" {
    var reader = std.Io.Reader.fixed("hello world\n");

    const s = count(&reader);
    try std.testing.expectEqual(s, 2);
}
