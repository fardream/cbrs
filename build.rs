use fefix::{codegen, Dictionary};
use std::{
    fs::{read_to_string, File},
    io::Write,
    path::PathBuf,
};

fn read_write(dict_name: &str, out_name: &str) -> Result<(), Box<dyn std::error::Error>> {
    let cargo_dir = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
    let dict_path = cargo_dir.join("fix-dicts").join(dict_name);

    let dict_content = read_to_string(dict_path)?;

    let dict = Dictionary::from_quickfix_spec(dict_content).unwrap();

    let def = codegen::gen_definitions(dict, &codegen::Settings::default());

    let out_path = cargo_dir.join("src").join(out_name);

    let mut out_file = File::create(out_path)?;

    out_file.write_all(def.as_bytes())?;

    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // read_write("sandbox.xml", "sandbox.rs")?;
    // read_write("production.xml", "production.rs")?;
    Ok(())
}
