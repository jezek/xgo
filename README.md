This is a wrapper for [github.com/jezek/xgb](https://github.com/jezek/xgb), which is a patched fork of [github.com/BurntSushi/xgb](https://github.com/BurntSushi/xgb).

The purpose of this package is to make window creation and window handling easier and adds keyboard typing and mouse moving abilities using X11 testing module.

# Dependencies
- [github.com/jezek/xgb](https://github.com/jezek/xgb)

# Konwn issues
- Keyboard writing doesn't write not mapped characters, if ```audacious```'s ```Global Hotkeys``` plugin is enabled.
