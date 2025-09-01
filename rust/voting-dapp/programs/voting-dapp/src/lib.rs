use anchor_lang::prelude::{borsh::de, *};

declare_id!("EmfKYKULNaYrApj3xJ7BD4zVG3Pfq5axGTeyPBSfjx2c");

#[program]
pub mod voting_dapp {
    use super::*;

    pub fn initialize_poll(ctx: Context<Initialize>, id: u64) -> Result<()> {
        msg!("Greetings from: {:?} {}", ctx.program_id, id);
        Ok(())
    }
}

#[account]
#[derive(InitSpace)]
pub struct Poll {
    pub id: u64,
}

#[derive(Accounts)]
#[instruction(poll_id: u64)]
pub struct Initialize<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
      init_if_needed,
      payer = signer,
      space = 8 + Poll::INIT_SPACE,
      seeds = [poll_id.to_le_bytes().as_ref()],
      bump,
    )]
    pub poll: Account<'info, Poll>,

    pub system_program: Program<'info, System>,
}

#[account]
#[derive(InitSpace)]
pub struct Candidate {
    pub id: u64,
}
