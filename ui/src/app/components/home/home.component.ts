import { Component, OnInit } from '@angular/core';
import { EnumService } from "../shared/services/Enum.service";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  inputs: ['Icons']
})
export class HomeComponent implements OnInit {
  svgList: string[] = [];
  socialMedia: ({ name: string; link: string; isCustom: boolean})[] = [];

  id = "tsparticles";
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
          mode: "push"
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

  ngOnInit(): void {
    this.svgList = Object.keys(this.enums.Icons);
    this.socialMedia = Object.values(this.enums.socialMedia);
  }

  constructor(public enums: EnumService) {

  }

}
