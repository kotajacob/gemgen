# gemgen

Command line tool for converting [Commonmark Markdown](https://commonmark.org/)
to [Gemtext](https://gemini.circumlunar.space/docs/gemtext.gmi). Gemgen uses the
[goldmark](https://pkg.go.dev/github.com/yuin/goldmark) markdown parser and [my
gemtext rendering module](https://git.sr.ht/~kota/goldmark-gemtext/).

The goal is to create proper _hand-made_ gemtext. Links and "autolinks" are
placed below each paragraph, but a "paragraph" of **only** links is left intact.
Normally, paragraphs are merged onto a single line, but hardlinks (double spaces
or \ at the end of a line) may be used for manual line breaks. Lists and
headings are simplified to the gemtext format, emphasis markings are removed (or
kept with the `-e` flag), horizontal rules are turned into 80 character lines,
and indented code is converted to the gemtext "fenced" format.

## Usage

```
gemgen [-e] [-i input.md] [-o output.gmi]
 -e : Keep emphasis symbols ( _ and * )
 -i : Read from a file instead of standard input.
 -o : Write to an output file instead of standard output.
```

## Example

Markdown
```md
> If there is a country that has committed unspeakable atrocities in the world,
> it is the United States of America. They don't care for human beings. - Nelson
> Mandela

In 1996, investigative journalist [Gary
Webb](https://en.wikipedia.org/wiki/Gary_Webb)
exposed a CIA-run business of selling cocaine produced in Nicaragua, to help
fund the anti-communist Contras in their fight against the Sandinistas in
Nicaragua. **These drugs were mostly sold to black communities in California, and
helped spark the Crack epidemic.** Several of the US dealers such as such as
[Ross](https://en.wikipedia.org/wiki/%22Freeway%22_Rick_Ross)
and [Oscar Danilo Blandon](https://en.wikipedia.org/wiki/Oscar_Danilo_Bland%C3%B3n),
were found to have CIA and DEA ties. Webb's reports were suppressed in the news
media. In 1997, Webb stated: "If we had met five years ago, you wouldn't have
found a more staunch defender of the newspaper industry than me ... And then I
wrote some stories that made me realize how sadly misplaced my bliss had been.
The reason I'd enjoyed such smooth sailing for so long hadn't been, as I'd
assumed, because I was careful and diligent and good at my job ... The truth was
that, in all those years, I hadn't written anything important enough to
suppress." In 2004, Webb was found dead in his home, shot in the back of the
head twice. His death was ruled a suicide.

_If We Must Die_ - Claude McKay

If we must die—let it not be like hogs\
Hunted and penned in an inglorious spot,\
While round us bark the mad and hungry dogs,\
Making their mock at our accursed lot.\
If we must die—oh, let us nobly die,\
So that our precious blood may not be shed\
In vain; then even the monsters we defy\
Shall be constrained to honor us though dead!\
Oh, Kinsmen! We must meet the common foe;\
Though far outnumbered, let us show us brave,\
And for their thousand blows deal one deathblow!\
What though before us lies the open grave?\
Like men we'll face the murderous, cowardly pack,\
Pressed to the wall, dying, but fighting back!\

[Claude McKay](https://poets.org/poet/claude-mckay)
[US Atrocities](https://github.com/dessalines/essays/blob/master/us_atrocities.md)
```

Gemtext
```gemtext
> If there is a country that has committed unspeakable atrocities in the world, it is the United States of America. They don't care for human beings. - Nelson Mandela

In 1996, investigative journalist Gary Webb exposed a CIA-run business of selling cocaine produced in Nicaragua, to help fund the anti-communist Contras in their fight against the Sandinistas in Nicaragua. These drugs were mostly sold to black communities in California, and helped spark the Crack epidemic. Several of the US dealers such as such as Ross and Oscar Danilo Blandon, were found to have CIA and DEA ties. Webb's reports were suppressed in the news media. In 1997, Webb stated: "If we had met five years ago, you wouldn't have found a more staunch defender of the newspaper industry than me ... And then I wrote some stories that made me realize how sadly misplaced my bliss had been. The reason I'd enjoyed such smooth sailing for so long hadn't been, as I'd assumed, because I was careful and diligent and good at my job ... The truth was that, in all those years, I hadn't written anything important enough to suppress." In 2004, Webb was found dead in his home, shot in the back of the head twice. His death was ruled a suicide.

=> https://en.wikipedia.org/wiki/Gary_Webb Gary Webb
=> https://en.wikipedia.org/wiki/%22Freeway%22_Rick_Ross Ross
=> https://en.wikipedia.org/wiki/Oscar_Danilo_Bland%C3%B3n Oscar Danilo Blandon

If We Must Die - Claude McKay

If we must die—let it not be like hogs
Hunted and penned in an inglorious spot,
While round us bark the mad and hungry dogs,
Making their mock at our accursed lot.
If we must die—oh, let us nobly die,
So that our precious blood may not be shed
In vain; then even the monsters we defy
Shall be constrained to honor us though dead!
Oh, Kinsmen! We must meet the common foe;
Though far outnumbered, let us show us brave,
And for their thousand blows deal one deathblow!
What though before us lies the open grave?
Like men we'll face the murderous, cowardly pack,
Pressed to the wall, dying, but fighting back!\

=> https://poets.org/poet/claude-mckay Claude McKay
=> https://github.com/dessalines/essays/blob/master/us_atrocities.md US Atrocities
```
