const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    var buffer: [10]u8 = undefined;

    std.debug.print("Let's play Guess the Number.\n", .{});
    std.debug.print("Pick a difficulty level (1, 2, or 3): ", .{});
    var input = try stdin.readUntilDelimiter(buffer[0..], '\n');

    const difficulty_level: i32 = switch (std.fmt.parseInt(i32, input, 10) catch 0) {
        1 => 10,
        2 => 100,
        3 => 1_000,
        else => {
            std.debug.print("Invalid option!!", .{});
            std.process.exit(0);
        },
    };

    const secret = std.crypto.random.intRangeAtMost(i32, 1, difficulty_level);
    var total_guesses: i32 = 1;
    std.debug.print("I have my number. What's your guess? ", .{});

    while (true) {
        input = try stdin.readUntilDelimiter(buffer[0..], '\n');
        const user_guess = std.fmt.parseInt(i32, input, 10) catch 0;

        if (user_guess > secret) {
            std.debug.print("Too high, Guess again: ", .{});
        } else if (user_guess < secret) {
            std.debug.print("Too low, Guess again: ", .{});
        } else {
            std.debug.print("You got it in {d} guesses!\n", .{total_guesses});
            break;
        }

        total_guesses += 1;
    }
}
