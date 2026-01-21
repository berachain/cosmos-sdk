# CodeQL Workflows

## Status

The advanced CodeQL workflows have been disabled (renamed with `.disabled` extension) due to conflicts with GitHub's default CodeQL setup.

### Disabled Workflows

- `codeql.yml.disabled` - Scheduled CodeQL analysis (Wednesday 7:23 AM) for multiple languages
- `codeqls.yml.disabled` - Scheduled CodeQL analysis (Friday 4:22 AM) for multiple languages
- `codeql-analysis.yml.disabled` - CodeQL analysis on PR/push for Go files

### Reason

GitHub's default CodeQL setup (configured via repository settings) cannot coexist with advanced CodeQL workflows. When both are enabled, you get this error:

```
Code Scanning could not process the submitted SARIF file:
CodeQL analyses from advanced configurations cannot be processed when the default setup is enabled
```

### Solution

**Option 1: Use Default Setup (Current)**
- Keep the default CodeQL setup enabled in repository settings
- Keep the advanced workflow files disabled (`.disabled` extension)
- GitHub will automatically scan the repository using default configurations

**Option 2: Use Advanced Workflows**
1. Disable the default CodeQL setup in repository settings:
   - Go to Settings → Code security and analysis
   - Under "Code scanning", disable the default setup
2. Re-enable the advanced workflows:
   ```bash
   mv .github/workflows/codeql.yml.disabled .github/workflows/codeql.yml
   mv .github/workflows/codeqls.yml.disabled .github/workflows/codeqls.yml
   mv .github/workflows/codeql-analysis.yml.disabled .github/workflows/codeql-analysis.yml
   ```
3. Commit and push the changes

### References

- [CodeQL Default Setup](https://docs.github.com/en/code-security/code-scanning/enabling-code-scanning/configuring-default-setup-for-code-scanning)
- [CodeQL Advanced Setup](https://docs.github.com/en/code-security/code-scanning/creating-an-advanced-setup-for-code-scanning/configuring-advanced-setup-for-code-scanning)
