package src

/*
# targets : 15.1-STABLE-HBSD
* x86-64-v2
* x86-64-v3

# src job ctx : hardened/15-stable/main | quarterly/hardened/15-stable/main-2026q4
* target?
* __FreeBSD_version => sys/sys/param.h
* __HardenedBSD_version =>  sys/sys/pax.h
* KERNEL = KERNCONF
* make.conf
* src.conf
* repo?
* CCACHE
* signing key?
* Integriforce ?

BUILD:
   - installer: *.iso, *.img
   - update
   - kernel.txz, base.txz, src.txz

*/

/*
   - type: git
     job: pkgbase
     git-repo-url: https://git.hardenedbsd.org/hardenedbsd/HardenedBSD
     git-branch: hardened/14-stable/master # new branch?
*/

/*
   pkgbase:
       vermeer: # hbsd-update -Bo ?
           - ["echo pkgbase vermeer test", "&&", "echo pkgbase vermeer test2"]

       goldmont-plus:
           - [
                 "echo pkgbase goldmont_plus test",
                 "&&",
                 "echo pkgbase goldmont_plus test2",
             ]
*/

// => https://github.com/kirby-101/HardenedBSD-src.git
type SrcJob struct{}

/*
 * lock
 * prepare
 * fetch src
 * build {installer,update,sets}
 * check
 * sign {installer,update,sets}
 * publish
 * cleanup
 */

/*
   "cleanuniverse",
   "universe",
   "universe-toolchain",
   "tinderbox",
   "worlds",
   "kernels",
   "kernel-toolchains",
   "targets",
   "toolchains",
   "makeman",
   "sysent",
*/
