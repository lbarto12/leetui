# AUR packaging

Reference copies of the PKGBUILDs published to the AUR. The AUR itself
is a separate git remote per package — these dirs are not used directly
by makepkg from this repo.

## Workflow per release

1. Cut a release on GitHub (push to `main` → `Auto-tag` workflow → `Release`
   workflow attaches `tar.gz` + `.sha256` assets).
2. For each AUR package, copy the matching PKGBUILD into a clone of the
   AUR repo and refresh it:

   ```sh
   cd /path/to/aur-leetui            # ssh://aur@aur.archlinux.org/leetui.git
   cp .../packaging/aur/leetui/PKGBUILD .
   updpkgsums                        # fills sha256sums
   makepkg --printsrcinfo > .SRCINFO
   git commit -am "v0.1.0"
   git push
   ```

   Repeat for `leetui-bin` and `leetui-git` (the `-git` PKGBUILD doesn't
   need version bumps — `pkgver()` derives it from the live repo).

## First-time AUR setup

```sh
ssh aur@aur.archlinux.org   # registers your SSH key on first try
git clone ssh://aur@aur.archlinux.org/leetui.git
git clone ssh://aur@aur.archlinux.org/leetui-bin.git
git clone ssh://aur@aur.archlinux.org/leetui-git.git
```
