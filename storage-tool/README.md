# Get Started

The Sirius Storage Tool is a user desktop app to interact with Sirius Storage.

## Download
[Windows](https://files.proximax.io/storage-tool/windows-storage-tool.zip)

[Linux](https://files.proximax.io/storage-tool/linux-storage-tool.tar.xz)

## Installation

### Windows
1. Open `file explorer` and find the `windows-storage-tool.zip` zipped file
2. To unzip the file, double-click the zipped folder to open it. Then, drag or copy the item from the zipped folder to a new location
3. Go to the new location and enter the `windows-storage-tool` folder.
4. Double click on `StorageClientApp.exe` to run it.

**Note**: You may be prompted by Windows that Microsoft Defender SmartScreen is preventing this unrecognized app from Starting.  Click on `More Info` and click on the `Run Anyway` button.

### Ubuntu Linux
Run the following commands in terminal

```bash
wget https://files.proximax.io/storage-tool/linux-storage-tool.tar.xz
tar -xvf linux-storage-tool.tar.xz
cd linux-storage-tool
./StorageClientApp
```

## Initial set up

1. Enter `Account name` and `Private Key`.  Click on `Generate Key` if you wish to use a new private key.  Please ensure the account has `XPX`.  For testnet2, you can get `XPX` from the [faucet](https://bctestnet2faucet.xpxsirius.io)

2. Click on `Save` to continue.

3. Enter the following information for settings:

- REST Server Address: `13.235.254.237:3000`

- Replicator Bootstrap Address: `13.213.19.108:7904` OR `18.142.186.205:3000`
- Local UDP Port: `6846` (Default Port)
- Account Name: Select the `Account Name` from Step 1
- Download folder: Click `Choose Directory` to select the directory you want to download files to from drives
- Fee Multiplier: `15000` (default fee multiplier for *Testnet2*)

4. Click on `Save` to save your configurations.

## Create drive

1. Click the `+` button in `Drives` tab to add drive.

2. Enter the following information for Create Drive prompt:

- **Name:** Name of the drive
- **Replicator Number:** Enter the number of replicators (minimum of 1 and maximum of 4)
- **Max Drive Size:** Enter a size below 500 MB
- **Local Drive Folder:** Click `Choose Directory` to select the directory you want the local drive to be located

## Delete drive

1. From the dropdown on the left panel in `Drives` tab, select the name of the drive to be deleted.

2. Click the `-` button in `Drives` tab to delete drive.

3. Notification will be received if drive is removed successfully.

## Upload files to replicators

1. When the drive is created, your files in local drive will appear in green color at the right half of the window. The right half represents difference between your `Local Drive` and `Replicated Drive`.

2. Click `Apply Changes` to upload files to replicators.

3. Notification will be received if the modification is registered and completed.

4. Your files will appear at the left half of the window, indicating that the `Replicated Drive` now contains them.

## Create Download channel

1. Click the `+` button in `Downloads` tab to add drive.

2. Enter the following information for Add new download channel prompt:

- **Name*:** Name of the download channel
- **Drive*:** Use `Select from my drives` dropdown to select the drive to download files from
- **Keys:** (Optional)
- **Prepaid*:** 1000

3. Notification will be received if download channel is created successfully.

## Download files

1. In the `Downloads` tab select the download channels to download files from from the dropdown menu next to Channels

2. In the `Downloads` tab on the left panel select the files to be downloaded by ticking the checkboxes.

3. Click the `>>` button in the middle to download the selected files.

4. Open the Download folder configured in iniitial set up to view the downloaded files.

## Delete Download channel

1. In the `Downloads` tab select the download channel to delete from from the dropdown menu next to Channels.

2. Click the `-` button in `Downloads` tab.

3. Press the `confirm` button in the confirmation prompt.

4. Notification will be received if download channel is created successfully.


