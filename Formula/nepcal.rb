class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/srishanbhattarai/nepcal/releases/download/v0.3.1/nepcal_0.3.1_darwin_amd64.tar.gz"
  version "0.3.1"
  sha256 "afa7266bc216bae9820175b146f20b8fee0fb80655419f1f9316712afacd0a04"

  def install
    bin.install "nepcal"
  end
end
