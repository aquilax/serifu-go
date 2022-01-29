package serifu_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aquilax/serifu-go"
)

const spec = `
# PAGE 1
- 1.1
Contract Text:/= The undersigned* agrees to sell his soul** for a thousand berries.***=/
Sign:/=
Menu:
- Pizza: 50 Yen
- Okonomiyaki: 100 Yen
- Beer: 200 Yen
=/
- 1.2
Menelaus/Announcing: Here in the mountains of Japan…
Menelaus/Announcing: …There is a steel cage made for one purpose.
- 1.3
- 1.4
Menelaus/Announcing:
Menelaus/Announcing:
Title: Moriking
Chapter Title: Chapter 31: Giant Asian Hornet vs. Palawan Stag Beetle
Shoko/Shadowed: ?!
- 1.5
* gasp (haa)
Shota/Sharp: A _death match?!?_ The invitation said it was gonna be arm wrestling...!
Menelaus/Announcing: It was changed at the last minute...
Menelaus/Announcing: ...at the strong insistence of the seeded contestant.
- 1.6
Palawan/Serious: I have no interest in such **pathetic games.**

# PAGE 2
- 2.1
Palawan/Serious: The only creatures with any right to live...
Palawan/Serious: ...are those with the beauty of strength.
* ha ha ha
- 2.2
Shota/Scared: The Palawan...
Shota/Scared: ...Stag Beetle...
Shoko/Bold: The what now?
- 2.3
Shota/Sharp: A giant stag beetle that lives on the Palawan archipelago in the Philippines!!
Shota/Sharp: With its overwhelming prowess in battle, it's said to be the strongest stag beetle on the planet!!
Shoko/Thought: Okay, so it's another cool bug, got it.

# PAGE 3
- 3.1
Palawan/Serious: You all disgust me.
- 3.2
Palawan: It is we Insecters who are the rightful masters of all life.
Palawan: And only the most powerful among us...
Palawan: ...is fit to rule the planet.
- 3.3
Palawan: Filthy pests and minor species from irrelevant islands...
Palawan: My world has no need for such trash.
- 3.4
Palawan: Send out your champion...
* glare (jiii)
Palawan: ...and I will end them.
- 3.5
Shota: What should we do? The only other battle to the death we did was with...
Shoko: Huh? Speaking of which, where's Oga?
Ko/Bold: Actually, I haven't seen him for a few days...!
Shoko: Yeah, he wasn't here for round two, either. Weird.
- 3.6
Oki: Ha ha ha, guess he got freaked out and split!
Oki: That's okay, I got this one!
Shoko/Thought: Wait, didn't Oki only get his butt whacked?
- 3.7
Moriking: I will face him.
Shoko/Bold: Mori--!
Meo/Bold: Hold up.
`

func ExampleParse() {
	got, _ := serifu.Parse(strings.NewReader(spec))
	b, _ := json.MarshalIndent(got, "", "  ")
	fmt.Println(string(b))
	// Output:
	// {
	//   "Pages": [
	//     {
	//       "Title": "PAGE 1",
	//       "IsSpread": false,
	//       "Panels": [
	//         {
	//           "Id": "1.1",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Contract Text",
	//               "Style": "",
	//               "Content": " The undersigned* agrees to sell his soul** for a thousand berries.**"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Sign",
	//               "Style": "",
	//               "Content": "Menu:\n- Pizza: 50 Yen\n- Okonomiyaki: 100 Yen\n- Beer: 200 Yen\n"
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "1.2",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Menelaus",
	//               "Style": "Announcing",
	//               "Content": "Here in the mountains of Japan…"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Menelaus",
	//               "Style": "Announcing",
	//               "Content": "…There is a steel cage made for one purpose."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "1.3",
	//           "Items": null
	//         },
	//         {
	//           "Id": "1.4",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Menelaus",
	//               "Style": "Announcing",
	//               "Content": ""
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Menelaus",
	//               "Style": "Announcing",
	//               "Content": ""
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Title",
	//               "Style": "",
	//               "Content": "Moriking"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Chapter Title",
	//               "Style": "",
	//               "Content": "Chapter 31: Giant Asian Hornet vs. Palawan Stag Beetle"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "Shadowed",
	//               "Content": "?!"
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "1.5",
	//           "Items": [
	//             {
	//               "Type": "soundEffect",
	//               "Name": "gasp ",
	//               "Transliteration": "(ha"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shota",
	//               "Style": "Sharp",
	//               "Content": "A _death match?!?_ The invitation said it was gonna be arm wrestling...!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Menelaus",
	//               "Style": "Announcing",
	//               "Content": "It was changed at the last minute..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Menelaus",
	//               "Style": "Announcing",
	//               "Content": "...at the strong insistence of the seeded contestant."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "1.6",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "Serious",
	//               "Content": "I have no interest in such **pathetic games.**"
	//             }
	//           ]
	//         }
	//       ]
	//     },
	//     {
	//       "Title": "PAGE 2",
	//       "IsSpread": false,
	//       "Panels": [
	//         {
	//           "Id": "2.1",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "Serious",
	//               "Content": "The only creatures with any right to live..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "Serious",
	//               "Content": "...are those with the beauty of strength."
	//             },
	//             {
	//               "Type": "soundEffect",
	//               "Name": "ha ha ha",
	//               "Transliteration": ""
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "2.2",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Shota",
	//               "Style": "Scared",
	//               "Content": "The Palawan..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shota",
	//               "Style": "Scared",
	//               "Content": "...Stag Beetle..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "Bold",
	//               "Content": "The what now?"
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "2.3",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Shota",
	//               "Style": "Sharp",
	//               "Content": "A giant stag beetle that lives on the Palawan archipelago in the Philippines!!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shota",
	//               "Style": "Sharp",
	//               "Content": "With its overwhelming prowess in battle, it's said to be the strongest stag beetle on the planet!!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "Thought",
	//               "Content": "Okay, so it's another cool bug, got it."
	//             }
	//           ]
	//         }
	//       ]
	//     },
	//     {
	//       "Title": "PAGE 3",
	//       "IsSpread": false,
	//       "Panels": [
	//         {
	//           "Id": "3.1",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "Serious",
	//               "Content": "You all disgust me."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "3.2",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "It is we Insecters who are the rightful masters of all life."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "And only the most powerful among us..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "...is fit to rule the planet."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "3.3",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "Filthy pests and minor species from irrelevant islands..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "My world has no need for such trash."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "3.4",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "Send out your champion..."
	//             },
	//             {
	//               "Type": "soundEffect",
	//               "Name": "glare ",
	//               "Transliteration": "(jii"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Palawan",
	//               "Style": "",
	//               "Content": "...and I will end them."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "3.5",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Shota",
	//               "Style": "",
	//               "Content": "What should we do? The only other battle to the death we did was with..."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "",
	//               "Content": "Huh? Speaking of which, where's Oga?"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Ko",
	//               "Style": "Bold",
	//               "Content": "Actually, I haven't seen him for a few days...!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "",
	//               "Content": "Yeah, he wasn't here for round two, either. Weird."
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "3.6",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Oki",
	//               "Style": "",
	//               "Content": "Ha ha ha, guess he got freaked out and split!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Oki",
	//               "Style": "",
	//               "Content": "That's okay, I got this one!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "Thought",
	//               "Content": "Wait, didn't Oki only get his butt whacked?"
	//             }
	//           ]
	//         },
	//         {
	//           "Id": "3.7",
	//           "Items": [
	//             {
	//               "Type": "text",
	//               "Source": "Moriking",
	//               "Style": "",
	//               "Content": "I will face him."
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Shoko",
	//               "Style": "Bold",
	//               "Content": "Mori--!"
	//             },
	//             {
	//               "Type": "text",
	//               "Source": "Meo",
	//               "Style": "Bold",
	//               "Content": "Hold up."
	//             }
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }

}
