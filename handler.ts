import { App } from '@slack/bolt';

const app = new App({
  signingSecret: process.env.SLACK_SIGNING_SECRET!,
  clientId: process.env.SLACK_CLIENT_ID,
  clientSecret: process.env.SLACK_CLIENT_SECRET,
  stateSecret: 'my-state-secret',
  scopes: ['commands', 'chat:write'],
  installationStore: {
    storeInstallation: async installation => {
      // TODO: 実際のデータベースに保存するために、ここを変更
      token_database[installation.team.id] = installation;
      return Promise.resolve();
    },
    fetchInstallation: async installQuery => {
      // TODO: 実際のデータベースから取得するために、ここを変更
      const installation = token_database[installQuery.teamId];
      return Promise.resolve(installation);
    },
  },
});
