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

## PAGE 2
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
! he is not really serious
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
	got, err := serifu.Parse(strings.NewReader(spec))
	b, _ := json.MarshalIndent(got, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	// Output:
	// {
	//   "pages": [
	//     {
	//       "title": "PAGE 1",
	//       "is_spread": false,
	//       "panels": [
	//         {
	//           "id": "1.1",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Contract Text",
	//               "style": "",
	//               "is_pre_formatted": true,
	//               "content": " The undersigned* agrees to sell his soul** for a thousand berries.**"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Sign",
	//               "style": "",
	//               "is_pre_formatted": true,
	//               "content": "Menu:\n- Pizza: 50 Yen\n- Okonomiyaki: 100 Yen\n- Beer: 200 Yen\n"
	//             }
	//           ]
	//         },
	//         {
	//           "id": "1.2",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Menelaus",
	//               "style": "Announcing",
	//               "is_pre_formatted": false,
	//               "content": "Here in the mountains of Japan…"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Menelaus",
	//               "style": "Announcing",
	//               "is_pre_formatted": false,
	//               "content": "…There is a steel cage made for one purpose."
	//             }
	//           ]
	//         },
	//         {
	//           "id": "1.3",
	//           "items": null
	//         },
	//         {
	//           "id": "1.4",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Menelaus",
	//               "style": "Announcing",
	//               "is_pre_formatted": false,
	//               "content": ""
	//             },
	//             {
	//               "type": "text",
	//               "source": "Menelaus",
	//               "style": "Announcing",
	//               "is_pre_formatted": false,
	//               "content": ""
	//             },
	//             {
	//               "type": "text",
	//               "source": "Title",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Moriking"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Chapter Title",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Chapter 31: Giant Asian Hornet vs. Palawan Stag Beetle"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "Shadowed",
	//               "is_pre_formatted": false,
	//               "content": "?!"
	//             }
	//           ]
	//         },
	//         {
	//           "id": "1.5",
	//           "items": [
	//             {
	//               "type": "soundEffect",
	//               "name": "gasp",
	//               "transliteration": "ha"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shota",
	//               "style": "Sharp",
	//               "is_pre_formatted": false,
	//               "content": "A _death match?!?_ The invitation said it was gonna be arm wrestling...!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Menelaus",
	//               "style": "Announcing",
	//               "is_pre_formatted": false,
	//               "content": "It was changed at the last minute..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Menelaus",
	//               "style": "Announcing",
	//               "is_pre_formatted": false,
	//               "content": "...at the strong insistence of the seeded contestant."
	//             }
	//           ]
	//         },
	//         {
	//           "id": "1.6",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "Serious",
	//               "is_pre_formatted": false,
	//               "content": "I have no interest in such **pathetic games.**"
	//             }
	//           ]
	//         }
	//       ]
	//     },
	//     {
	//       "title": "PAGE 2",
	//       "is_spread": true,
	//       "panels": [
	//         {
	//           "id": "2.1",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "Serious",
	//               "is_pre_formatted": false,
	//               "content": "The only creatures with any right to live..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "Serious",
	//               "is_pre_formatted": false,
	//               "content": "...are those with the beauty of strength."
	//             },
	//             {
	//               "type": "soundEffect",
	//               "name": "ha ha ha",
	//               "transliteration": ""
	//             }
	//           ]
	//         },
	//         {
	//           "id": "2.2",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Shota",
	//               "style": "Scared",
	//               "is_pre_formatted": false,
	//               "content": "The Palawan..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shota",
	//               "style": "Scared",
	//               "is_pre_formatted": false,
	//               "content": "...Stag Beetle..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "Bold",
	//               "is_pre_formatted": false,
	//               "content": "The what now?"
	//             }
	//           ]
	//         },
	//         {
	//           "id": "2.3",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Shota",
	//               "style": "Sharp",
	//               "is_pre_formatted": false,
	//               "content": "A giant stag beetle that lives on the Palawan archipelago in the Philippines!!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shota",
	//               "style": "Sharp",
	//               "is_pre_formatted": false,
	//               "content": "With its overwhelming prowess in battle, it's said to be the strongest stag beetle on the planet!!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "Thought",
	//               "is_pre_formatted": false,
	//               "content": "Okay, so it's another cool bug, got it."
	//             }
	//           ]
	//         }
	//       ]
	//     },
	//     {
	//       "title": "PAGE 3",
	//       "is_spread": false,
	//       "panels": [
	//         {
	//           "id": "3.1",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "Serious",
	//               "is_pre_formatted": false,
	//               "content": "You all disgust me."
	//             },
	//             {
	//               "type": "sideNote",
	//               "content": "he is not really serious"
	//             }
	//           ]
	//         },
	//         {
	//           "id": "3.2",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "It is we Insecters who are the rightful masters of all life."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "And only the most powerful among us..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "...is fit to rule the planet."
	//             }
	//           ]
	//         },
	//         {
	//           "id": "3.3",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Filthy pests and minor species from irrelevant islands..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "My world has no need for such trash."
	//             }
	//           ]
	//         },
	//         {
	//           "id": "3.4",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Send out your champion..."
	//             },
	//             {
	//               "type": "soundEffect",
	//               "name": "glare",
	//               "transliteration": "jii"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Palawan",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "...and I will end them."
	//             }
	//           ]
	//         },
	//         {
	//           "id": "3.5",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Shota",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "What should we do? The only other battle to the death we did was with..."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Huh? Speaking of which, where's Oga?"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Ko",
	//               "style": "Bold",
	//               "is_pre_formatted": false,
	//               "content": "Actually, I haven't seen him for a few days...!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Yeah, he wasn't here for round two, either. Weird."
	//             }
	//           ]
	//         },
	//         {
	//           "id": "3.6",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Oki",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "Ha ha ha, guess he got freaked out and split!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Oki",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "That's okay, I got this one!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "Thought",
	//               "is_pre_formatted": false,
	//               "content": "Wait, didn't Oki only get his butt whacked?"
	//             }
	//           ]
	//         },
	//         {
	//           "id": "3.7",
	//           "items": [
	//             {
	//               "type": "text",
	//               "source": "Moriking",
	//               "style": "",
	//               "is_pre_formatted": false,
	//               "content": "I will face him."
	//             },
	//             {
	//               "type": "text",
	//               "source": "Shoko",
	//               "style": "Bold",
	//               "is_pre_formatted": false,
	//               "content": "Mori--!"
	//             },
	//             {
	//               "type": "text",
	//               "source": "Meo",
	//               "style": "Bold",
	//               "is_pre_formatted": false,
	//               "content": "Hold up."
	//             }
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParseString() {
	got, err := serifu.Parse(strings.NewReader(spec))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(got)

	// Output:
	// # PAGE 1
	//     - 1.1
	//     Contract Text:/= The undersigned* agrees to sell his soul** for a thousand berries.**=/
	//     Sign:/=Menu:
	// - Pizza: 50 Yen
	// - Okonomiyaki: 100 Yen
	// - Beer: 200 Yen
	// =/

	//     - 1.2
	//     Menelaus/Menelaus: Here in the mountains of Japan…
	//     Menelaus/Menelaus: …There is a steel cage made for one purpose.

	//     - 1.3

	//     - 1.4
	//     Menelaus/Menelaus:
	//     Menelaus/Menelaus:
	//     Title: Moriking
	//     Chapter Title: Chapter 31: Giant Asian Hornet vs. Palawan Stag Beetle
	//     Shoko/Shoko: ?!

	//     - 1.5
	//     * gasp (ha)
	//     Shota/Shota: A _death match?!?_ The invitation said it was gonna be arm wrestling...!
	//     Menelaus/Menelaus: It was changed at the last minute...
	//     Menelaus/Menelaus: ...at the strong insistence of the seeded contestant.

	//     - 1.6
	//     Palawan/Palawan: I have no interest in such **pathetic games.**

	// ## PAGE 2
	//     - 2.1
	//     Palawan/Palawan: The only creatures with any right to live...
	//     Palawan/Palawan: ...are those with the beauty of strength.
	//     * ha ha ha

	//     - 2.2
	//     Shota/Shota: The Palawan...
	//     Shota/Shota: ...Stag Beetle...
	//     Shoko/Shoko: The what now?

	//     - 2.3
	//     Shota/Shota: A giant stag beetle that lives on the Palawan archipelago in the Philippines!!
	//     Shota/Shota: With its overwhelming prowess in battle, it's said to be the strongest stag beetle on the planet!!
	//     Shoko/Shoko: Okay, so it's another cool bug, got it.

	// # PAGE 3
	//     - 3.1
	//     Palawan/Palawan: You all disgust me.
	//     ! he is not really serious

	//     - 3.2
	//     Palawan: It is we Insecters who are the rightful masters of all life.
	//     Palawan: And only the most powerful among us...
	//     Palawan: ...is fit to rule the planet.

	//     - 3.3
	//     Palawan: Filthy pests and minor species from irrelevant islands...
	//     Palawan: My world has no need for such trash.

	//     - 3.4
	//     Palawan: Send out your champion...
	//     * glare (jii)
	//     Palawan: ...and I will end them.

	//     - 3.5
	//     Shota: What should we do? The only other battle to the death we did was with...
	//     Shoko: Huh? Speaking of which, where's Oga?
	//     Ko/Ko: Actually, I haven't seen him for a few days...!
	//     Shoko: Yeah, he wasn't here for round two, either. Weird.

	//     - 3.6
	//     Oki: Ha ha ha, guess he got freaked out and split!
	//     Oki: That's okay, I got this one!
	//     Shoko/Shoko: Wait, didn't Oki only get his butt whacked?

	//     - 3.7
	//     Moriking: I will face him.
	//     Shoko/Shoko: Mori--!
	//     Meo/Meo: Hold up.
}
