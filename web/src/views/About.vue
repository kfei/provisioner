<template>
  <v-container class="pa-5">
    <v-layout>
      <v-flex xs12>
        <div class="title mb-1 grey--text">
          Question
        </div>
        <div class="body-2">
          Please build a skeleton / poc of real-time coupon provisioning web application.
        </div>
        <div class="title mt-4 mb-1 grey--text">
          Requirement
        </div>
        <div class="body-2">
          1. User Story -- What user can do only via the web app are: (1) See the number of coupons left; (2) Submit a request to get a coupon.

          2. Design -- There will be just a number ("Coupons Available") and a button ("Give Me A Coupon") on the page. The fundamental frame/layout will be a navigation bar, a side menu containing 2 menu items and 2 content subpages which one of them will place the number and button, the other subpage can be blank. The number is telling people how many available coupons left. The button is letting user to claim a coupon.

          3. Backlog -- The number of coupons left will start from 0 and increase by 10 every 1 second after server's lanuching. For example, when server starts, the number is 0; 1 second later, the number is 10; 2 seconds later, the number is 20; ... never stops.

          4. Backlog -- The number of coupons will be deducted by N after user clicking the button "Give Me A Coupon" for N times. For example, server is kept running for 4 seconds, the number of coupons left is 40. Now user clicks the button once, the number of coupons left changes to 39. 1 second later, the number of coupons left added up to 49 because the coupons keep increasing by 10 every 1 second. If user clicks the button twice, the number of coupons left will be deducted by 2. If user clicks the button 10 times, the number will decrease by 10.

          5. JS -- Please consider using React.js + React Router to build front end components, Redux to handle data from/to backend, websocket to get the real-time coupons data and submit "Give Me A Coupon" request.

          6. CSS -- Provide some basic css styles to decorate the nav bar and 2 side menus, the nav bar stays at the top and side menu stays at the left, on the right is the subpage of each menu item.

          7. Backend -- Backend web server can be any language/framework.

          8. No Auth -- User doesn't need to login. Server doesn't need to know who gets how many coupons.

          9. Simple Data Storage -- The number of coupons left can be in any kind of store: db,file,cache.

          10. Test -- Please consider adding one unit test case.

          11. Output -- Dependency/Package manager and build tool can be anything you feel comfortable with. Should be a deployable application and the source code.
        </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<style lang="stylus" scoped>
.body-2
  white-space pre-line
</style>
