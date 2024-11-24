FROM node:latest AS base
WORKDIR /build

FROM base AS install
RUN mkdir -p /temp/deps
COPY package.json /temp/deps
RUN cd /temp/deps && npm install

# copy node_modules from temp directory
# then copy all (non-ignored) project files into the image
FROM base AS prerelease
COPY --from=install /temp/deps/node_modules node_modules
COPY . .
ENV NODE_ENV=production
RUN npm run build-only

# copy production dependencies and source code into final image
FROM nginx:stable-alpine AS release
COPY --from=prerelease /build/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=prerelease /build/dist /usr/share/nginx/html/
