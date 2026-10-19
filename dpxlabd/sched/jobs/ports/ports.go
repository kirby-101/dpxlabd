package ports

/*
# targets : 15.1-STABLE-HBSD
* x86-64-v2
* x86-64-v3

# port job ctx : 15.1-STABLE-HBSD
* target?
* __FreeBSD_version ?
* repo?
* package list?
* overlays?
* make.conf
* CCACHE
* signing key?
*/
/*
poudriere bulk -v -j hbsd-14-stable -p default -O hyprland_overlay -f /usr/local/etc/poudriere.d/packages-default

poudriere sets für desktop und laptop
*/

// => https://github.com/kirby-101/HardenedBSD-ports.git
type PortsJob struct{}

/*
   - type: git
     job: ports
     git-repo-url: https://git.hardenedbsd.org/hardenedbsd/ports
     git-branch: hardenedbsd/main
*/

/*
ports:
        goldmont-plus:
            - [
                  "poudriere ports -u -p default -p hyprland_overlay -p signal_desktop_overlay",
                  "&&",
                  "poudriere bulk -v -O hyprland_overlay -O signal_desktop_overlay -j HBSD-14-STABLE -p default -z goldmont_plus -f /usr/local/etc/poudriere.d/pkglist",
                  "&&",
                  "gotify push -t goldmont-plus -p 10 'Build erfolgreich'",
                  "||",
                  "gotify push -t goldmont-plus -p 9 'Build fehlgeschlagen'",
              ]

        vermeer:
            - [
                  "poudriere ports -u -p default -p hyprland_overlay -p signal_desktop_overlay",
                  "&&",
                  "poudriere bulk -v -O hyprland_overlay -O signal_desktop_overlay -j HBSD-14-STABLE -p default -z vermeer -c -f /usr/local/etc/poudriere.d/pkglist",
                  "&&",
                  "gotify push -t vermeer -p 10 'Build erfolgreich'",
                  "||",
                  "gotify push -t vermeer -p 9 'Build fehlgeschlagen'",
              ]
*/
