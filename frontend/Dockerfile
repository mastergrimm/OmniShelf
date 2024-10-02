FROM node:20-alpine

WORKDIR /frontend

COPY ./package.json .
COPY ./pnpm-lock.yaml .

RUN npm i -g pnpm
RUN pnpm i

EXPOSE 5173

CMD ["pnpm", "run", "dev", "--", "--host", "0.0.0.0"]
