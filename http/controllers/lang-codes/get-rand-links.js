const fs = require('fs')
const axios = require('axios')
const cheerio = require('cheerio')

const unparsedContent = fs.readFileSync('./langguages.json');
const parsedContent = JSON.parse(unparsedContent)

const withRandLinksExecution = parsedContent.Merged.map(async language => {
	const { data } = await axios.get(`https://${language.LangCode}.wikipedia.org/`)

	const $ = cheerio.load(data)
	const randLink = $('li#n-randompage a').attr('href')

	return Object.assign(language, {RandLink: randLink})
})

Promise.all(withRandLinksExecution).then(withRandLinks => {
	fs.writeFileSync('./languages-with-random.json', JSON.stringify(withRandLinks))
})