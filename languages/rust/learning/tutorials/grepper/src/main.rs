fn main() {
    let args = args().collect::<Vec<String>>();
    with(run(args), |it| match it {
        Ok(()) => exit(0),
        Err(err) => {
            eprintln!("{}: {}", style_error("Problem encountered"), err);
            exit(1);
        }
    });
}

fn run(args: Vec<String>) -> Result<(), Box<dyn Error>> {
    match is_stdin_piped() {
        true => piped_grep(PipedGrepOptionsBuilder::parse(args)?)?,
        false => grep(GrepOptionsBuilder::parse(args)?)?,
    }
    Ok(())
}
