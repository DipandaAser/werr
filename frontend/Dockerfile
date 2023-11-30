# build stage
FROM node:20.9 as build
LABEL dev="Aser <aserdipanda@gmail.com>"

# RUN npm install -g pnpm

WORKDIR /app
# copy all the projet config files
COPY . .
# install dependencies
RUN npm install

# build the SvelteKit app
RUN npm run build

# run stage, to separate it from the build stage, to save disk storage y
FROM node:20.9

WORKDIR /app 

# copy stuff from the build stage
COPY --from=build /app/package*.json ./
COPY --from=build /app/build ./

# expose the app's port
EXPOSE 3000
# run the server
CMD ["node", "-r", "dotenv/config", "./index.js"]