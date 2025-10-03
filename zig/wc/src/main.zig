const std = @import("std");

pub fn main() !void {
    var buffer: [1024]u8 = undefined;
    var stdin = std.fs.File.stdin().reader(&buffer);
    const reader = &stdin.interface;

    var count_lines = false;

    var args = std.process.args();
    _ = args.skip();

    while (args.next()) |arg| {
        if (std.mem.eql(u8, arg, "-l")) {
            count_lines = true;
        }
    }

    const sum = count(reader, count_lines);
    std.debug.print("sum: {d}\n", .{sum});
}

fn count(reader: *std.Io.Reader, count_lines: bool) i32 {
    const contents = reader.allocRemaining(std.heap.page_allocator, .unlimited) catch return 0;

    var it: std.mem.TokenIterator(u8, .any) = undefined;
    var wc: i32 = 0;

    if (count_lines) {
        it = std.mem.tokenizeAny(u8, contents, "\n");
    } else {
        it = std.mem.tokenizeAny(u8, contents, " \n\t\r");
    }

    while (it.next()) |_| {
        wc += 1;
    }

    return wc;
}

test "test count words" {
    var reader = std.Io.Reader.fixed("word1 word2 word3 word4");

    const s = count(&reader, false);
    try std.testing.expectEqual(s, 4);
}

test "test count lines" {
    var reader = std.Io.Reader.fixed("word1 word2 word3\nline2\nline3 word1");

    const s = count(&reader, true);
    try std.testing.expectEqual(s, 3);
}
