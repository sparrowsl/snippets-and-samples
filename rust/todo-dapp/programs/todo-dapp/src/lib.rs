use anchor_lang::prelude::*;

declare_id!("2pkC5ywNgXcT1cSdTcQbF9eYKLHi9ZqmBysRMaWZXyC2");

#[program]
pub mod todo_dapp {
    use super::*;

    pub fn initialize_todo(ctx: Context<Initialize>, id: u64, task: String) -> Result<()> {
        msg!("Greetings from: {:?}", ctx.program_id);
        let todo = &mut ctx.accounts.todo;
        todo.id = id;
        todo.task = task;
        todo.is_done = false;

        Ok(())
    }

    pub fn update_todo(ctx: Context<UpdateTodo>, task: String, is_done: bool) -> Result<()> {
        let todo = &mut ctx.accounts.todo;
        msg!("Greetings!! updating todo with id: {}", todo.id);
        todo.task = task;
        todo.is_done = is_done;

        Ok(())
    }

    pub fn delete_todo(ctx: Context<DeleteTodo>) -> Result<()> {
        msg!(
            "Greetings!! deleting todo with id: {}",
            ctx.accounts.todo.id
        );

        Ok(())
    }
}

#[account]
#[derive(InitSpace)]
pub struct Todo {
    pub id: u64,

    #[max_len(200)]
    pub task: String,

    pub is_done: bool,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
      init_if_needed,
      payer = signer,
      space = 8 + Todo::INIT_SPACE,
      seeds = [b"todo", signer.key().as_ref()],
      bump,
    )]
    pub todo: Account<'info, Todo>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdateTodo<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
      mut,
      seeds = [b"todo", signer.key().as_ref()],
      bump,
    )]
    pub todo: Account<'info, Todo>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct DeleteTodo<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
      mut,
      close = signer,
      seeds = [b"todo", signer.key().as_ref()],
      bump,
    )]
    pub todo: Account<'info, Todo>,

    pub system_program: Program<'info, System>,
}
