---
title: Validação de CPF e CNPJ
tags:
    - c_sharp
    - cpf
    - cnpj
---

```csharp
public static class ValidateDocument
{
    ///
    /// Validates CPF/CNPJ and reformats any missing left padding '0'
    ///
    public static bool FormatIfValidDocument(
        string document, 
        out string formattedAccount)
    {
        document = string.Join("", document
            .Where(c => c is >= '0' and <= '9')
            .Select(c => c));
        if (IsValidCpf(document) || IsValidCnpj(document))
        {
            formattedAccount = document;
            return true;
        }

        // Try CPF
        var cpfDoc = document.PadLeft(11, '0');
        if (IsValidCpf(cpfDoc))
        {
            formattedAccount = cpfDoc;
            return true;
        }

        // Try CNPJ
        var cnpjDoc = document.PadLeft(14, '0');
        if (IsValidCnpj(cnpjDoc))
        {
            formattedAccount = cnpjDoc;
            return true;
        }

        formattedAccount = "";
        return false;
    }

    private static readonly byte[] CpfMult1 = { 10, 9, 8, 7, 6, 5, 4, 3, 2 };
    private static readonly byte[] CpfMult2 = { 11, 10, 9, 8, 7, 6, 5, 4, 3, 2 };
    private static readonly byte[] CnpjMult1 = { 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 };
    private static readonly byte[] CnpjMult2 = { 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 };

    private static bool IsValidCpf(string cpf) => IsValidNumber(cpf, CpfMult1, CpfMult2);
    
    private static bool IsValidCnpj(string cnpj) => IsValidNumber(cnpj, CnpjMult1, CnpjMult2);


    private static bool IsValidNumber(string doc, IReadOnlyCollection<byte> mult1, IReadOnlyCollection<byte> mult2)
    {
        byte GetModule11(IReadOnlyList<byte> data, IEnumerable<byte> mult)
        {
            var sum = mult.Select((t, i) => data[i] * t).Sum();
            var rest = sum % 11;
            return (byte)(rest < 2 ? 0 : (11 - rest));
        }

        var size = mult2.Count + 1;
        if (doc.Length != size || doc.All(c => c == doc[0]))
            // Test if length is OK and all numbers are not equal
            return false;
        var digits = doc.Select(d => (byte)(d - '0')).ToArray();

        var rest1 = GetModule11(digits, mult1);
        var updatedDigits = digits[..mult1.Count].Append(rest1).ToArray();
        var rest2 = GetModule11(updatedDigits, mult2);

        return doc.EndsWith($"{rest1}{rest2}");
    }
}
```