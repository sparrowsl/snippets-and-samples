use anchor_lang::prelude::*;

declare_id!("8A98FQ61VnEQFh5F1pkCMagNaQuuv7Rd3udNGMPggH5m");

#[program]
pub mod url_shortener {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, id: u64, original: String) -> Result<()> {
        msg!("Greetings from: {:?}", ctx.program_id);
        let link = &mut ctx.accounts.link;
        link.id = id;
        link.original_url = original;
        link.short_url = generate_short_url();
        Ok(())
    }

    pub fn create_link(ctx: Context<UpdateLink>, original: String) -> Result<()> {
        msg!("Greetings from: {:?}", ctx.program_id);
        let link = &mut ctx.accounts.link;
        link.original_url = original;
        link.short_url = generate_short_url();
        Ok(())
    }

    pub fn delete_link(_ctx: Context<DeleteLink>) -> Result<()> {
        Ok(())
    }
}

#[account]
#[derive(InitSpace)]
pub struct Shortener {
    id: u64,
    #[max_len(255)]
    original_url: String,
    #[max_len(100)]
    short_url: String,
}

// generate a random string generator link
fn generate_short_url() -> String {
    use std::collections::hash_map::DefaultHasher;
    use std::hash::{Hash, Hasher};

    // Generate a simple hash-based short code
    let mut hasher = DefaultHasher::new();
    let timestamp = std::time::SystemTime::now()
        .duration_since(std::time::UNIX_EPOCH)
        .unwrap_or_default()
        .as_nanos();

    timestamp.hash(&mut hasher);
    let hash = hasher.finish();

    // Convert to base36 for URL-friendly characters (0-9, a-z)
    let mut result = String::new();
    let mut num = hash;
    let chars = "0123456789abcdefghijklmnopqrstuvwxyz";

    // Generate 6-character short code
    for _ in 0..6 {
        let idx = (num % 36) as usize;
        result.push(chars.chars().nth(idx).unwrap());
        num /= 36;
    }

    return result;
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
      init_if_needed,
      payer = signer,
      space = 8 + Shortener::INIT_SPACE,
      seeds = [b"shortener", signer.key().as_ref()],
      bump,
    )]
    pub link: Account<'info, Shortener>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdateLink<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
        mut,
        seeds = [b"shortener", signer.key().as_ref()],
        bump,
    )]
    pub link: Account<'info, Shortener>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct DeleteLink<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,

    #[account(
        mut,
        close = signer,
        seeds = [b"shortener", signer.key().as_ref()],
        bump,
    )]
    pub link: Account<'info, Shortener>,

    pub system_program: Program<'info, System>,
}
