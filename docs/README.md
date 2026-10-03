# Appleseed

Appleseed generates Apple-style more typeable strong passwords.

+ Apple-style?
  This type of password, e.g., `funrus-Hommez-kajzo7`, was developed by Apple in an effort to encourage good password practices among users and developers.
  For more about the Apple background, see [this blog post][mondelloBlog] and [this conference video][mondelloVideo], both by Ricky Mondello.
+ More typeable?
  These passwords have several features that make them relatively easy to type on all sorts of keyboards.
  In particular, they have only one digit, one uppercase letter, and two hyphens.
  Everything else in the 20-character password is a lowercase letter.
+ Strong?
  Yes.
  Like Apple's, these passwords have about [71 bits of entropy][entropy].
  In addition, [zxcvbn][zxcvbn] typically gives them a 4/4 score and estimates that they would take centuries to crack.

[mondelloBlog]: https://rmondello.com/2024/10/07/apple-passwords-generated-strong-password-format
[mondelloVideo]: https://youtu.be/-0dwX2kf6Oc
[entropy]: https://support.apple.com/guide/security/automatic-strong-passwords-secc84c811c4/web
[zxcvbn]: https://lowe.github.io/tryzxcvbn/

## Usage

Appleseed exports two functions.

```go
// The password will be, e.g., "funrus-Hommez-kajzo7".
pw := appleseed.Password()

// The password will be, e.g., "funrusHommezkajzo7".
pw := appleseed.PasswordWithoutHyphens()
```

## Format

I haven't found a full description of the format from Apple.
There is [a Javascript implementation][appleJS], but it may not be authoritative.
Ricky Mondello, who worked on these passwords at Apple, [gave a talk][mondelloVideo] and [wrote a blog post][mondelloBlog] about the format.
Putting this all together, here is my understanding of these passwords.

Keep this example in mind: `funrus-Hommez-kajzo7`.

+ The full password is twenty characters; seventeen are letters, one is a digit, and two are hyphens.
+ One of the seventeen letters is uppercase; the rest are lowercase.
  The consonants are `bcdfghjkmnpqrstvwxz` and `BCDFGHJKMNPQRSTVWXZ`.
  Note: `l` and `L` are not included.
  The vowels are `aeiouy` and `AEIUY`.
  Note: `o` is included, but not `O`.
+ The hyphens divide the password into three six-character quasi-words.
  More precisely, the hyphens must appear as the seventh and fourteenth characters in the password.
+ The digit is placed in one of five positions: before one of the two hyphens, after one of the two hyphens, or at the very end of the password.
  More precisely, the digit must appear as the sixth, eighth, thirteenth, fifteenth, or twentieth character in the password.
+ Other than the digit, the characters run largely in sequences of consonant-vowel-consonant.

I assume that they omit some letters to avoid confusion.
Uppercase `O` is easily confused with `0` in some fonts.
Lowercase `l` is easily confused with `1` or `I` in some fonts.

On the other hand, Apple may omit `L` only as an implementation detail.
[Apple's Javascript implementation][appleJS] classifies `L` as an "unambiguous" character.
However, `L` never appears in a password since the implementation uses `toUpperCase` on a lowercase letter, and `l` is never used.
But it isn't clear to me that they think that `L` itself is a problem.

[appleJS]: https://developer.apple.com/password-rules/scripts/generator.js

## Differences from Apple

Appleseed has one significant difference from Apple-generated passwords.
A randomly generated string may contain offensive substrings.
Apple avoids returning offensive passwords by keeping a list of offensive substrings and rejecting passwords that contain them.
Appleseed does not have such a list, and it doesn't screen the random passwords it generates.
