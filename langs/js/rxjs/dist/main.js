"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const rxjs_1 = require("rxjs");
const fs_1 = require("fs");
const readline_1 = __importDefault(require("readline"));
function main() {
    return __awaiter(this, void 0, void 0, function* () {
        console.time('exec-time');
        const rl = readline_1.default.createInterface({
            input: (0, fs_1.createReadStream)('../../../data-generator/short-product-data.ndjson'),
            terminal: true
        });
        yield new Promise((resolve) => {
            (0, rxjs_1.fromEventPattern)((handler) => rl.on('line', handler), (handler) => rl.on('end', handler))
                .pipe((0, rxjs_1.filter)((chunk) => typeof chunk === 'string'), (0, rxjs_1.map)((line) => JSON.parse(line)), (0, rxjs_1.mergeMap)((product) => (0, rxjs_1.from)(product.Prices)), (0, rxjs_1.map)((price) => price.Country), (0, rxjs_1.reduce)((countryCount, country) => {
                var _a;
                countryCount[country] = ((_a = countryCount[country]) !== null && _a !== void 0 ? _a : 0) + 1;
                return countryCount;
            }, {}))
                .subscribe((result) => {
                console.log(result);
            })
                .add(resolve);
        });
        console.timeEnd('exec-time');
    });
}
main().catch(console.error);
