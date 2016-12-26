import Ember from 'ember';

const { service } = Ember.inject;

export default Ember.Route.extend({
  session: Ember.inject.service('session'),

  model(params) {
      const session = this.get('session');

      if (session.get('isAuthenticated')) {
          session.set('skipRedirectOnInvalidation', true);
          session.invalidate();
      }

      this.get('store').query('user', {
          token: params.token
      }).then(users => {
          this.controllerFor("signup").set('signuptoken', params.token);
          this.controllerFor("signup").set('signupemail', users.get("firstObject").get('email'))
      });
  }
});
