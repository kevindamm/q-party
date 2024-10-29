<template>
  <div class="about">
    <h1>About this Game</h1>

    <p>
      I made this as a private service to have fun and practice, and to keep my
      aging brain accessing minutiae regularly.  I've done my best to keep any
      hints of the original piece of art (Jeopardy!) separated from the game
      implementation.  While the github repo does maintain an index of the games
      as recorded on <a href="https://j-archive.com/">J!-Archive</a> and the
      game server does retain a copy of those questions, I don't make any of the
      questions or answers (or definitions and questions) publicly available.
    </p>

    <p>
      If you are interested in hosting your own board inventions on this server,
      reach out to me via github or email or socials.  If you want to run your
      own instance of the game, I have instructions for hosting your own version
      of the server on the github repo but I will not help setting up on any
      systems other than the one described.  I only have so much time, you
      understand.
    </p>

    <p>
      Another reason I'm keeping this private is that I am not dedicating much
      attention to its maintenance or improvement.  However, I'm considering
      adding voice recognition for player answering, something I've had some
      interest in building and have a solid plan for, but consider it too big
      a distraction.  If I knew there were many others <i>also</i> interested
      in that feature, it would be the one significant improvement I'd add. <br>
      Maybe I should set up a donation button for it?
    </p>

    <p>
      I have many more ongoing projects that I'm excited to get to share soon!
      Let me know if you have a game you are especially fond of and it could get
      bumped up on my list :smile:  A little motivation goes a long way.
    </p>
  </div>
</template>

<style>
.about h1 {
  text-align: center;
  text-shadow: 2px 2px 3px;
}

@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
