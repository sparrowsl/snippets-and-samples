const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [20]u8 = undefined;

    const Item = struct { price: f32, quantity: i32 };

    var items: [3]Item = undefined;

    for (0..items.len) |idx| {
        std.debug.print("Enter the price of item {d}: ", .{idx + 1});
        var input = try stdin.readUntilDelimiter(buffer[0..], '\n');
        const price = try std.fmt.parseFloat(f32, input);

        std.debug.print("Enter the quantity of item {d}: ", .{idx + 1});
        input = try stdin.readUntilDelimiter(buffer[0..], '\n');
        const quantity = try std.fmt.parseInt(i32, input, 10);

        items[idx] = Item{ .price = price, .quantity = quantity };
    }

    var subtotal: f32 = 0;
    for (0..items.len) |idx| {
        subtotal = subtotal + (items[idx].price * @as(f32, @floatFromInt(items[idx].quantity)));
    }

    const tax_rate: f32 = 5.5;
    const tax = (tax_rate / 100) * subtotal;

    std.debug.print("Subtotal: ${d:.2}\n", .{subtotal});
    std.debug.print("Tax: ${d:.2}\n", .{tax});
    std.debug.print("Total: ${d:.2}\n", .{subtotal + tax});
}
