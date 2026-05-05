# AUR packaging

These PKGBUILDs are **templates**. The release workflow renders them with
the real version + sha256 sums and pushes them to AUR automatically on
every push to `main`. You should not need to edit them by hand.

The packages:

- `leetui` — source build of the latest tagged release
- `leetui-bin` — prebuilt binary of the latest tagged release
- `leetui-git` — VCS package; builds from `main` HEAD

## One-time setup

1. Create an AUR account at https://aur.archlinux.org/register and add an
   SSH public key under your account settings.
2. On your dev machine, confirm SSH works: `ssh aur@aur.archlinux.org` should
   greet you and exit.
3. Add the matching **private** key to this GitHub repo's secrets as
   `AUR_SSH_PRIVATE_KEY` (Settings → Secrets and variables → Actions →
   New repository secret). Paste the file contents including the
   `-----BEGIN ... PRIVATE KEY-----` lines.

That's it. The first push to `main` after this triggers a release and
populates all three AUR packages.

## Manual edits

If you need to change a PKGBUILD itself (e.g. add a dependency), edit
the file in `packaging/aur/<pkg>/PKGBUILD` here, push to main, and the
workflow handles the rest.
