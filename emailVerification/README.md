# Domain Record Checker

This Go application checks the given domain for the following records:

- **MX (Mail Exchange)**: Determines if the domain has mail exchange records.
- **SPF (Sender Policy Framework)**: Checks for SPF records that define authorized mail servers.
- **DMARC (Domain-based Message Authentication, Reporting & Conformance)**: Verifies DMARC records for email authentication policies.

## Prerequisites

Ensure you have the following installed:

- **Go (latest version)** - [Download Go](https://golang.org/dl/)

## Installation

Clone the repository or create a new directory and add the code manually:

```bash
git clone <repository-url>
cd <repository-folder>

```

Build the application:

```bash
go build -o domain-checker
```

## usage

Run the application:

```bash
./domain-checker
```

Enter domain names one by one and press Enter. Example:

```bash
example.com
google.com
```

To exit, press Ctrl + D (Linux/Mac) or Ctrl + Z (Windows).

## Output Format

The program prints domain records in the following format:

```bash
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
example.com, true, true, v=spf1 include:_spf.example.com ~all, true, v=DMARC1; p=none;
```


### Instructions for Use
1. Replace `<repository-url>` with the actual URL of your repository.
2. Save this content in a file named `README.md` in the root of your project directory.
3. Open the file in Visual Studio Code to view the formatted README.
