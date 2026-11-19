package main

var iconsMap = map[string]string{
	"Simbolo-vegano-300x300":               "🌱",
	"Origem-animal-site":                   "🥩",
	"Gluten-site":                          "🌾",
	"Leite-e-derivados-site":               "🥛",
	"Ovo-site":                             "🍳",
	"Alergenicos-site":                     "⚠️",
	"Simbolo-pimenta-300x300":              "🌶️",
	"vegan":                                "🌱",
	"carne":                                "🥩",
	"gluten":                               "🌾",
	"lactose":                              "🥛",
	"ovo":                                  "🍳",
	"476bf979-2cbb-476b-8739-02ed26485235": "🐷",
	"daf04cd5-bacd-4ea4-91ce-48ea45cb0ac4": "⚠️",
	"Indicado para veganos":                "🌱",
	"Contém pimenta":                       "🌶️",
	"Contêm produtos de origem suína":      "🐷",
	"Contém ingrediente(s) potencialmente alergênico(s)": "⚠️",
	"Contêm produtos de origem animal":                   "🥩",
	"Contém glúten":                                      "🌾",
	"Contém leite e/ou derivados":                        "🥛",
	"Contêm ovos":                                        "🍳",
}

var mealOrder = []string{"breakfast", "lunch", "dinner"}

var mealTypeHeaders = map[string]string{
	"breakfast": "*CAFÉ DA MANHÃ*",
	"lunch":     "*ALMOÇO*",
	"dinner":    "*JANTAR*",
}

var weekDay = []string{
	"DOMINGO",
	"SEGUNDA",
	"TERÇA",
	"QUARTA",
	"QUINTA",
	"SEXTA",
	"SÁBADO",
}

const legend = `🌱 - Indicado para veganos
🥩 - Contém produtos de origem animal
🐷 - Contém produtos de origem suína
🌾 - Não indicado para celíacos por conter glúten
🥛 - Não indicado para intolerantes à lactose por conter lactose
🍳 - Contêm ovos
⚠️ - Contém produto(s) alergênico(s)
🍯 - Contém mel
🌶️ - Contém pimenta`

const title = `*CARDÁPIO RU %s - %s - %s*`

const warningAI = "Estimativa de calorias gerada por IA com base em uma porção individual. Pode conter imprecisões."

const taken = `Cardápio obtido de forma automatizada pelo site oficial do restaurante universitário, disponível no link: %s.`

const channel = `Esta mensagem foi enviada para o canal do WhatsApp, disponível no link: %s.`

const copyright = "Esta mensagem e este canal não possuem relação com a universidade nem com o restaurante universitário."
