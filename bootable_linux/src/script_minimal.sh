# check if the caller has root permissions
SUDO_VAR=""
if [ "$EUID" != 0 ];
then
    echo "EUID is not 0"
    SUDO_VAR="sudo "
else
    echo "EUID is 0"
fi

echo "Starting the script"

# clean working directory
$SUDO_VAR rm -r working_dir

# install dependencies
$SUDO_VAR apt update
$SUDO_VAR apt install -y git cpio
$SUDO_VAR apt install -y build-essential flex libncurses5-dev bc libelf-dev bison libssl-dev
$SUDO_VAR apt install -y qemu qemu-system

# Create a working directory
mkdir working_dir
cd working_dir

# clone linux repo and build
git clone --depth=1 git://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git
cd linux
make x86_64_defconfig
make -j $(nproc)
cd ../

# create a custom init file and build with static linking
gcc -static -o init ../dummy_init.c
chmod +x init

# create initramfs directory hierarchy
mkdir initramfs
cp init initramfs/

# create initramfs cpio file
cd initramfs
find . | cpio -H newc -o > ../initramfs.cpio
cd ..

# launch qemu
qemu-system-x86_64 -kernel linux/arch/x86_64/boot/bzImage -nographic -append "console=ttyS0" -initrd initramfs.cpio -m 1G
