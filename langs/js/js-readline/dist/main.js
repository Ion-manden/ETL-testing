"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = require("fs");
const readline_1 = __importDefault(require("readline"));
const countryCount = {};
readline_1.default.createInterface({
    input: (0, fs_1.createReadStream)('../../../data-generator/product-data.ndjson'),
    terminal: false
}).on('line', function (line) {
    var _a;
    const product = JSON.parse(line);
    const prices = product.Prices;
    for (const price of prices) {
        const country = price.Country;
        countryCount[country] = ((_a = countryCount[country]) !== null && _a !== void 0 ? _a : 0) + 1;
    }
}).on('close', () => {
    console.log(countryCount);
});
