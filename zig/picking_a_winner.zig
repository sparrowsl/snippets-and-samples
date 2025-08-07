const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [100]u8 = undefined;

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var names = std.ArrayList([]const u8).init(allocator);

    while (true) {
        std.debug.print("Enter a name: ", .{});
        const input = try stdin.readUntilDelimiter(buffer[0..], '\n');
        const trimmed = std.mem.trim(u8, input, " ");

        if (std.mem.eql(u8, trimmed, "")) break;

        try names.append(input);
    }

    if (names.items.len == 0) std.process.exit(0);
    _ = pick_winner(names);
}

fn pick_winner(list: std.ArrayList([]const u8)) []const u8 {
    var prng = std.Random.DefaultPrng.init(blk: {
        var seed: u64 = undefined;
        std.posix.getrandom(std.mem.asBytes(&seed)) catch {};
        break :blk seed;
    });

    if (list.items.len == 1) return list.items[0];

    const random = prng.random();
    const secret_number = random.intRangeLessThan(u64, 0, std.math.cast(u64, list.items.len) orelse 0);

    return list.items[secret_number];
}
