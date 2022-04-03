import { Injectable } from '@angular/core';
import { Icons } from '../../../models/Icon.enum';


@Injectable({providedIn: 'root'})
export class EnumService {
    Icons = Icons;
    socialMedia = {
        linkedIn: {name: 'linkedIn', link:'https://www.linkedin.com/in/nicholas-figgins/', isCustom: true},
        github: {name: 'github', link: 'https://github.com/nickfiggins', isCustom: true},
        //resume: {name: 'library_books', link: 'assets/files/resume.pdf', isCustom: false},
        buymeacoffee: {name: 'buymeacoffee', link: 'https://www.buymeacoffee.com/nickfiggins', isCustom: true},
        gopeer: {name: 'perm_contact_calendar', link: 'https://gopeer.org/profile/603fc26b7ec0544baa4304c2', isCustom: false},
        ycbm: {name: 'ycbm', link: 'https://nickfiggins.youcanbook.me/', isCustom: true}
    };
    driveFolders = {
        France: {
            name: 'France',
            id: '1F3qUFe-IzqMshCX3zLKZs1diZ_IiqWOH',
        },
        Prague: {
            name: 'Prague',
            id: '1--I5AfI8qIm5dQ1yS_bycZ-TtTOMy3qe'
        }
    }
  particlesOptions = {
    background: {
      color: {
        value: "#000000"
      },
      position: "50% 50%",
      repeat: "no-repeat",
      size: "cover"
    },
    fullScreen: {
      enable: true,
      zIndex: 1
    },
    interactivity: {
      events: {
        onClick: {
          enable: true,
          mode: "repulse"
        },
        onHover: {
          enable: true,
          mode: "bubble"
        }
      },
      modes: {
        bubble: {
          distance: 400,
          duration: 0.3,
          opacity: 1,
          size: 4
        },
        grab: {
          distance: 400,
          links: {
            opacity: 0.5
          }
        }
      }
    },
    particles: {
      links: {
        color: {
          value: "#ffffff"
        },
        distance: 500,
        opacity: 0.4,
        width: 2
      },
      move: {
        attract: {
          rotate: {
            x: 600,
            y: 1200
          }
        },
        direction: "bottom",
        enable: true,
        outModes: {
          bottom: "out",
          left: "out",
          right: "out",
          top: "out"
        }
      },
      number: {
        density: {
          enable: true
        },
        value: 400
      },
      opacity: {
        random: {
          enable: true
        },
        value: 0.5,
        animation: {
          minimumValue: 0.1,
          speed: 1
        }
      },
      size: {
        random: {
          enable: true
        },
        value: 10,
        animation: {
          minimumValue: 0.1,
          speed: 40
        }
      }
    }
};
}