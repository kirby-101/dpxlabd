// Pseudo-Code mit argon2-browser
async function solve(challenge, difficulty) {
  const [saltB64, expires] = challenge.split(":");
  const salt = Uint8Array.from(atob(saltB64), c => c.charCodeAt(0));

  let nonce = 0;
  while (true) {
    const input = new TextEncoder().encode(saltB64 + ":" + nonce);
    const hash = await argon2.hash({
      pass: input,
      salt: salt,
      time: 1,
      mem: 16384,
      parallelism: 1,
      hashLen: 32,
      type: argon2.ArgonType.Argon2id
    });

    if (hash.hash[0] === 0 && hash.hash[1] === 0) { // difficulty 2
      return { nonce, hash: Array.from(hash.hash) };
    }
    nonce++;
  }
}
